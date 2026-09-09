package assemblyDataHoldStructures

import (
	"LDT/src/functions/conversion"
	"LDT/src/functions/manifestDownload"
	"LDT/src/functions/pkgInstallation"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type Package struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type InstalledAssembly struct {
	Name              string    `json:"name"`
	LocationPath      string    `json:"locationPath"`
	LCVersion         string    `json:"lcVersion"`
	InstalledPackages []Package `json:"installedPackages"`
}

func (a InstalledAssembly) GetCode() string {
	code := a.LCVersion
	for _, pkg := range a.InstalledPackages {
		code += fmt.Sprintf(":%s-%s", pkg.Author, pkg.Name)
	}
	return code
}

func (a InstalledAssembly) CreateFromCode(code string, locationPath string) {
	a.LocationPath = locationPath
	parts := strings.Split(code, ":")
	a.LCVersion = parts[0]
	a.InstalledPackages = []Package{}
	for _, part := range parts[1:] {
		pkgParts := strings.Split(part, "-")
		a.InstalledPackages = append(a.InstalledPackages, Package{
			Name:   pkgParts[1],
			Author: pkgParts[0],
		})
	}
	return
}

func (a InstalledAssembly) Install() error {
	versionToDownload, err := conversion.StrToLCVersion(a.LCVersion)
	if err != nil {
		return err
	}

	manifestID := versionToDownload.ManifestID
	installationPath := a.LocationPath
	err = manifestDownload.DownloadManifest(manifestID, installationPath)
	if err != nil {
		return err
	}

	for _, pkg := range a.InstalledPackages {
		link := fmt.Sprintf("https://thunderstore.io/package/%s/%s/", pkg.Author, pkg.Name)
		err = pkgInstallation.InstallPkgWithDependenciesByLCVersion(link, a.LocationPath, versionToDownload)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a InstalledAssembly) Launch() error {
	binaryPath := filepath.Join(a.LocationPath, fmt.Sprintf("%s %s", a.LCVersion, a.Name))
	cmd := exec.Command(binaryPath)

	//cmd.Stderr = os.Stderr //uncomment only when debugging

	return cmd.Start()
}
