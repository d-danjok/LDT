package structures

import (
	"LDT/src/download"
	"LDT/src/extraction"
	"LDT/src/manifestDownload"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type RemoteSrc struct {
	RemoteType string
	Url        string
}

type Assembly struct {
	Name                  string
	BaseVersionManifestID string
	Mods                  RemoteSrc
}

func (a *Assembly) SetMods(remoteType string, remoteLink string) {
	a.Mods.RemoteType = remoteType
	a.Mods.Url = remoteLink
}

func (a *Assembly) downloadAll(tmpDirPath string, steamFolder string) error {
	err := os.MkdirAll(tmpDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Printf("\n")

	//download base
	err = manifestDownload.DownloadManifest(
		a.BaseVersionManifestID,
		filepath.Join(steamFolder,
			"steamapps/content/app_1966720",
			a.Name,
		))

	if err != nil {
		return err
	}

	//download Mods archive
	err = downloads.DownloadByRemoteType(
		a.Mods.RemoteType,
		a.Mods.Url,
		tmpDirPath,
		"Mods.zip")
	if err != nil {
		return err
	}

	return nil
}

func (a *Assembly) Install(steamFolder string) error {
	//ensure installation folder will be opened when installation finishes
	defer func() {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("explorer", filepath.Join(steamFolder, "steamapps/content/app_1966720", a.Name))
			// Run the command
			_ = cmd.Start()

		}
	}()

	//download base version and Mods
	err := a.downloadAll(os.TempDir(), steamFolder)
	if err != nil {
		return err
	}

	if a.Mods.RemoteType == "" {
		return nil
	}

	//extract Mods to steam folder
	err = extraction.ExtractToSteamFolder(os.TempDir(), "Mods.zip", a.Name, steamFolder)
	if err != nil {
		return err
	}

	return nil
}
