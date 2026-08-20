package structures

import (
	"LCAD/src/download"
	"LCAD/src/extraction"
	"LCAD/src/manifestDownload"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type remoteSrc struct {
	remoteType string
	url        string
}

type Assembly struct {
	Name                  string
	BaseVersionManifestID string
	mods                  remoteSrc
}

func (a *Assembly) SetMods(remoteType string, remoteLink string) {
	a.mods.remoteType = remoteType
	a.mods.url = remoteLink
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

	//download mods archive
	err = downloads.DownloadByRemoteType(
		a.mods.remoteType,
		a.mods.url,
		tmpDirPath,
		"mods.zip")
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

	//download base version and mods
	err := a.downloadAll(os.TempDir(), steamFolder)
	if err != nil {
		return err
	}

	if a.mods.remoteType == "" {
		return nil
	}

	//extract mods to steam folder
	err = extraction.ExtractToSteamFolder(os.TempDir(), "mods.zip", a.Name, steamFolder)
	if err != nil {
		return err
	}

	return nil
}
