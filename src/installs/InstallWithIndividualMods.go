package installs

import (
	definitions "LCAD/src"
	"LCAD/src/cli"
	"LCAD/src/fsManagement"
	"LCAD/src/pkgInstallation"
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
)

var lcVersion string

func installNewLCInstance() (string, error) {
	var installationName string
	var installationPath string

	lcVersionToInstallNum, err := cli.SelectByNum("LC version to install", len(definitions.LCVersions), cli.ListLCVersions, nil)
	if err != nil {
		return "", err
	}

	normaliseLCVersionName := func(lcVersionName string) string {
		switch lcVersionName {
		case "v50HotFix":
			return "v50"
		default:
			return lcVersionName
		}
	}

	lcVersion = normaliseLCVersionName(definitions.LCVersions[lcVersionToInstallNum].Name)

	fmt.Printf("How do you want you installation to be named?\n")
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Printf("\n: ")
		scanner.Scan()
		installationName = scanner.Text()

		installationPath = filepath.Join(definitions.SteamFolder, "steamapps/content/app_1966720", lcVersion+" "+installationName)
		if !fsManagement.Exists(installationPath) {
			break
		}
		fmt.Printf("There is an assembly with the same name, select different name\n\n")
	}

	err = cli.LocateSteamFolder()
	if err != nil {
		return "", fmt.Errorf("error locating Steam folder: %v", err)
	}

	err = definitions.LCVersions[lcVersionToInstallNum].Install(installationPath)
	if err != nil {
		return "", err
	}

	return installationPath, nil
}

func installIntoExistingInstance() (string, error) {
	fmt.Printf("Locate existing LC installation root folder")
	installationPath, err := dialog.Directory().Title("Locate existing LC installation root folder").Browse()

	cli.ClearTerminal()

	if fsManagement.Exists(filepath.Join(installationPath, "BepInEx")) {
		fmt.Printf("LC installation you have located already has some mods installed\n")

		installedModsTreatmentType, err := cli.SelectByNum("new mod installation type",
			2,
			nil,
			[]string{"Install alongside", "Erase existing"},
		)
		if err != nil {
			return "", err
		}
		switch installedModsTreatmentType {
		case 0:
			pkgInstallation.BepInExInstalled = true
		case 1:
			err = os.RemoveAll(filepath.Join(installationPath, "BepInEx"))
			if err != nil {
				return "", err
			}
		}
	}
	cli.ClearTerminal()

	lcVersionToInstallNum, err := cli.SelectByNum("LC version you have", len(definitions.LCVersions), cli.ListLCVersions, nil)
	if err != nil {
		return "", err
	}

	normaliseLCVersionName := func(lcVersionName string) string {
		switch lcVersionName {
		case "v50HotFix":
			return "v50"
		default:
			return lcVersionName
		}
	}

	lcVersion = normaliseLCVersionName(definitions.LCVersions[lcVersionToInstallNum].Name)

	return installationPath, nil
}

func isThunderstoreLink(link string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "https" || parsedURL.Host != "thunderstore.io" {
		return false
	}

	if !strings.HasSuffix(parsedURL.Path, "/") {
		return false
	}

	parts := strings.Split(parsedURL.Path, "/")
	if (len(parts) == 7 && parts[6] != "") || !(len(parts) == 6 || len(parts) == 7) {
		return false
	}

	return parts[1] == "c" && parts[2] == "lethal-company" && parts[3] == "p" &&
		parts[4] != "" && parts[5] != ""
}

func installMods(lcLocationPath string) error {
	var link string

	defer cli.ClearTerminal()

	isValidLink := isThunderstoreLink

	for true {
		fmt.Printf("Enter link to mod you want to install from Thunderstore" +
			"link should be in following format: \n" +
			"\thttps://thunderstore.io/c/lethal-company/p/[author]/[packageName]/\n" +
			"\nIf you would like to abort mod installation type \"abort\"\n\n: ")
		_, err := fmt.Scanln(&link)
		if err != nil {
			return err
		}
		if link == "a" || link == "abort" {
			return nil
		}
		if !isValidLink(link) {
			fmt.Printf("%s is not a valid link\n", link)

			cli.ClearTerminal()
			continue
		}

		err = pkgInstallation.InstallPkgWithDependenciesByLCVersion(link, lcLocationPath, lcVersion)
		if err != nil {
			fmt.Printf("Failed to install mod ")
			if confirmation, err := cli.Confirm("continue"); confirmation && err != nil {

				cli.ClearTerminal()
				continue
			}
			return err
		}

		cli.ClearTerminal()

		fmt.Printf("Mod installed successfully\n")
		installNext, err := cli.Confirm("install the next mod")
		if err != nil {
			return err
		}
		if !installNext {
			return nil
		}
		
		cli.ClearTerminal()
	}

	return nil
}

func InstallWithIndividualMods() error {
	var installationPath string

	installationType, err := cli.SelectByNum("installation type",
		2,
		nil,
		[]string{"Install new LC instance", "Install into existing LC installation"})
	if err != nil {
		return err
	}

	cli.ClearTerminal()

	switch installationType {
	case 0:
		installationPath, err = installNewLCInstance()
	case 1:
		installationPath, err = installIntoExistingInstance()
	}
	if err != nil {
		return err
	}

	cli.ClearTerminal()

	err = installMods(installationPath)
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("explorer", installationPath)
		// Run the command
		_ = cmd.Start()

	}

	return nil
}
