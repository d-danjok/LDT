package pkgInstallation

import (
	downloads "LDT/src/download"
	"LDT/src/structures"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const maxInstallationTries = 5

type Version struct {
	VersionNumber string   `json:"version_number"`
	DownloadURL   string   `json:"download_url"`
	Description   string   `json:"description"`
	Dependencies  []string `json:"dependencies"`
	FullName      string   `json:"full_name"`
}

type VersionRelease struct {
	Version string
	Date    time.Time
}

// GetVersionDates scrapes the /versions page and returns Version→date pairs.
func GetVersionDates(author, pkgName string) ([]VersionRelease, error) {
	link := fmt.Sprintf(
		"https://thunderstore.io/c/lethal-company/p/%s/%s/versions/",
		author, pkgName,
	)

	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// The page has a table where each <tr> contains:
	//   <td><a href="...">VERSION</a></td>
	//   <td>DATE STRING</td>
	//   <td>DOWNLOADS</td>
	//   <td>...</td>
	var results []VersionRelease
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			cells := tdTexts(n)
			if len(cells) >= 2 {
				version := strings.TrimSpace(cells[0])
				dateStr := strings.TrimSpace(cells[1])
				if version != "" && dateStr != "" && version != "Version" {
					t, err := time.Parse("Jan 2, 2006, 3:04 PM", dateStr)
					if err == nil {
						results = append(results, VersionRelease{Version: version, Date: t})
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	return results, nil
}

// tdTexts extracts the text content of each <td> in a <tr>.
func tdTexts(tr *html.Node) []string {
	var cells []string
	for n := tr.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode && n.Data == "td" {
			cells = append(cells, innerText(n))
		}
	}
	return cells
}

// innerText recursively collects all text under a node.
func innerText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(innerText(c))
	}
	return sb.String()
}

func getPkg(author, name, version string) (Version, error) {
	link := fmt.Sprintf("https://thunderstore.io/api/experimental/package/%s/%s/%s/", author, name, version)

	resp, err := http.Get(link)
	if err != nil {
		return Version{}, err
	}
	defer resp.Body.Close()

	var ver Version
	if err := json.NewDecoder(resp.Body).Decode(&ver); err != nil {
		return Version{}, err
	}
	return ver, nil
}

func getMaxPkgDateByLCVersion(version structures.LCVersion) (time.Time, error) {
	var maxDate time.Time
	var err error

	formDate := func(date string) (time.Time, error) {
		maxDate, err = time.Parse(time.DateOnly, date)
		if err != nil {
			return time.Time{}, err
		}
		return maxDate, nil
	}

	return formDate(version.ReleaseDate)
}

// resolveAndInstallPkg finds the latest Version of a package released before
// maxDate, downloads it, installs it, and returns its dependencies.
func resolveAndInstallPkg(pkgAuthor, pkgName, destPath string, maxDate time.Time) ([]string, error) {
	fullName := pkgAuthor + "-" + pkgName

	pkgVersions, err := GetVersionDates(pkgAuthor, pkgName)
	if err != nil {
		return nil, fmt.Errorf("fetching versions for %s: %w", fullName, err)
	}

	var versionToDownload VersionRelease
	var lastVersionFound VersionRelease
	var versionWasFound bool
	for _, version := range pkgVersions {
		if version.Date.Before(maxDate) {
			versionWasFound = true
			versionToDownload = version
			break
		}
		lastVersionFound = version
	}
	if !versionWasFound {
		versionToDownload = lastVersionFound
	}

	var pkgInfo Version
	for i := 0; i < maxInstallationTries; i++ {
		fmt.Printf("Installing %s (Try number %d)\n", fullName, i+1)

		pkgInfo, err = getPkg(pkgAuthor, pkgName, versionToDownload.Version)
		if err != nil {
			err = fmt.Errorf("fetching package %s: %w", pkgName, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = downloads.DownloadFileByLink(pkgInfo.DownloadURL, filepath.Join(destPath, fullName+".zip"))
		if err != nil {
			return nil, fmt.Errorf("downloading %s: %w", fullName, err)
		}

		err = InstallPkg(fullName, destPath, destPath)
		if err != nil {
			err = fmt.Errorf("installing package %s: %w", pkgName, err)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}

	//remove archive
	err = os.Remove(filepath.Join(destPath, fullName+".zip"))
	if err != nil {
		return nil, err
	}

	//Avoiding ddos protection
	time.Sleep(200 * time.Millisecond)

	return pkgInfo.Dependencies, nil
}

func installDependencies(dependencies []string, destPath string, maxDate time.Time) error {
	for _, dependency := range dependencies {
		nameParts := strings.Split(dependency, "-")

		pkgAuthor := nameParts[0]
		pkgName := nameParts[1]

		deps, err := resolveAndInstallPkg(pkgAuthor, pkgName, destPath, maxDate)
		if err != nil {
			return err
		}

		if err = installDependencies(deps, destPath, maxDate); err != nil {
			return err
		}
	}

	return nil
}

func InstallPkgWithDependenciesByLCVersion(link string, destPath string, LCVersion structures.LCVersion) error {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return err
	}

	linkPathParts := strings.Split(parsedURL.Path, "/")

	partShift := 0
	if link[len(link)-1] == '/' {
		partShift = 1
	}

	pkgName := linkPathParts[len(linkPathParts)-1-partShift]
	pkgAuthor := linkPathParts[len(linkPathParts)-2-partShift]

	maxPkgDate, err := getMaxPkgDateByLCVersion(LCVersion)
	if err != nil {
		return err
	}

	deps, err := resolveAndInstallPkg(pkgAuthor, pkgName, destPath, maxPkgDate)
	if err != nil {
		return err
	}

	return installDependencies(deps, destPath, maxPkgDate)
}
