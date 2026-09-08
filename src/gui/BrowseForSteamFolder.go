package gui

import (
	"LDT/src/appdata/userdata"
	"fmt"

	"github.com/sqweek/dialog"
)

func BrowseForSteamFolder() error {
	var err error

	userdata.General.SteamFolderLocation, err = dialog.Directory().
		Title("Select steam folder").
		Browse()

	if err != nil {
		return err
	}
	fmt.Printf("\n Steam folder located successfully\n")

	a := userdata.General
	_ = a

	return nil
}
