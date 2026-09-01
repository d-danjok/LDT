package gui

import (
	definitions "LDT/src"
	"fmt"
	"path/filepath"

	"github.com/sqweek/dialog"
)

func BrowseForSteamFolder() error {
	var err error

	definitions.SteamFolder, err = dialog.Directory().
		Title("Select steam folder").
		SetStartDir(
			filepath.Join(
				definitions.SteamFolder,
				definitions.LCAssembliesDefFolderSubPath,
			)).
		Browse()

	if err != nil {
		return err
	}
	fmt.Printf("\n Steam folder located successfully\n")

	return nil
}
