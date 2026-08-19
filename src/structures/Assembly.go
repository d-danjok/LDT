package structures

import (
	"LCAD/src/download"
	"LCAD/src/extraction"
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
	Name string
	base remoteSrc
	mods remoteSrc
}

func (a *Assembly) SetBase(remoteType string, remoteLink string) {
	a.base.remoteType = remoteType
	a.base.url = remoteLink
}

func (a *Assembly) SetMods(remoteType string, remoteLink string) {
	a.mods.remoteType = remoteType
	a.mods.url = remoteLink
}

func (a *Assembly) downloadArchives(tmpDirPath string) error {
	err := os.MkdirAll(tmpDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	//download base archive
	err = downloads.DownloadByRemoteType(
		a.base.remoteType,
		a.base.url,
		tmpDirPath,
		"base.zip")
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
	//download archives with base version and mods
	err := a.downloadArchives(os.TempDir())
	if err != nil {
		return err
	}

	err = extraction.ExtractToSteamFolder(os.TempDir(), "base.zip", a.Name, steamFolder)
	if err != nil {
		return err
	}

	if a.mods.remoteType == "" {
		return nil
	}

	err = extraction.ExtractToSteamFolder(os.TempDir(), "mods.zip", a.Name, steamFolder)
	if err != nil {
		return err
	}

	//open installation folder if windows is used
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("explorer", filepath.Join(steamFolder, "steamapps/content/app_1966720", a.Name))
		// Run the command
		_ = cmd.Start()

	}

	return nil
}
