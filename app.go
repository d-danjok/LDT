package main

import (
	"context"
	definitions "LDT/src"
	"LDT/src/fsManagement"
	"LDT/src/installs"
	"LDT/src/pkgInstallation"
	"LDT/src/preloads"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"net/url"
	"strings"

	"github.com/sqweek/dialog"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	preloads.PreloadAssemblyList()
}

type AssemblyInfo struct {
	Name string `json:"name"`
}

type LCVersionInfo struct {
	Name     string `json:"name"`
	LastDate string `json:"lastDate"`
}

func (a *App) GetAssemblies() []AssemblyInfo {
	result := make([]AssemblyInfo, len(definitions.Assemblies))
	for i, asm := range definitions.Assemblies {
		result[i] = AssemblyInfo{Name: asm.Name}
	}
	return result
}

func (a *App) GetLCVersions() []LCVersionInfo {
	result := make([]LCVersionInfo, len(definitions.LCVersions))
	for i, v := range definitions.LCVersions {
		result[i] = LCVersionInfo{Name: v.Name, LastDate: v.LastDate}
	}
	return result
}

func (a *App) GetDefaultSteamFolder() string {
	if runtime.GOOS == "windows" {
		return "C:\\Program Files (x86)\\Steam"
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".steam", "steam")
}

func (a *App) BrowseForSteamFolder() (string, error) {
	path, err := dialog.Directory().Title("Locate Steam folder").Browse()
	if err != nil {
		return "", err
	}
	return path, nil
}

func (a *App) BrowseForLCInstallation() (string, error) {
	path, err := dialog.Directory().Title("Locate existing LC installation root folder").Browse()
	if err != nil {
		return "", err
	}
	return path, nil
}

func (a *App) InstallCompleteAssembly(assemblyIndex int, steamFolder string) error {
	if assemblyIndex < 0 || assemblyIndex >= len(definitions.Assemblies) {
		return fmt.Errorf("invalid assembly index")
	}
	definitions.SteamFolder = steamFolder
	assembly := &definitions.Assemblies[assemblyIndex]
	return assembly.Install(steamFolder)
}

func (a *App) InstallNewInstance(steamFolder string, lcVersionIndex int, instanceName string) (string, error) {
	if lcVersionIndex < 0 || lcVersionIndex >= len(definitions.LCVersions) {
		return "", fmt.Errorf("invalid LC version index")
	}
	definitions.SteamFolder = steamFolder
	lcVersion := definitions.LCVersions[lcVersionIndex]
	installationPath := filepath.Join(steamFolder, "steamapps/content/app_1966720", lcVersion.Name+" "+instanceName)
	if fsManagement.Exists(installationPath) {
		return "", fmt.Errorf("an installation with this name already exists")
	}
	err := lcVersion.Install(installationPath)
	if err != nil {
		return "", err
	}
	return installationPath, nil
}

func (a *App) CheckHasBepInEx(installationPath string) bool {
	return fsManagement.Exists(filepath.Join(installationPath, "BepInEx"))
}

func (a *App) SetExistingInstallation(installationPath string, lcVersionIndex int, eraseExisting bool) error {
	if lcVersionIndex < 0 || lcVersionIndex >= len(definitions.LCVersions) {
		return fmt.Errorf("invalid LC version index")
	}
	if eraseExisting {
		err := os.RemoveAll(filepath.Join(installationPath, "BepInEx"))
		if err != nil {
			return err
		}
	} else {
		pkgInstallation.BepInExInstalled = true
	}
	return nil
}

func (a *App) IsValidThunderstoreLink(link string) bool {
	return installs.IsThunderstoreLink(link)
}

func (a *App) InstallMod(link string, installationPath string, lcVersionIndex int) error {
	if lcVersionIndex < 0 || lcVersionIndex >= len(definitions.LCVersions) {
		return fmt.Errorf("invalid LC version index")
	}
	lcVersion := definitions.LCVersions[lcVersionIndex]
	return pkgInstallation.InstallPkgWithDependenciesByLCVersion(link, installationPath, lcVersion)
}

func (a *App) OpenFolder(path string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}
	_ = cmd.Start()
}

func (a *App) ParseThunderstoreLink(link string) (string, string, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", "", err
	}
	parts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(parts) < 4 {
		return "", "", fmt.Errorf("invalid link format")
	}
	author := parts[len(parts)-2]
	pkgName := parts[len(parts)-1]
	return author, pkgName, nil
}
