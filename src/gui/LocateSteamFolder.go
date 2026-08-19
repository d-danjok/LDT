package gui

import (
	definitions "LCAD/src"
	"fmt"

	"github.com/sqweek/dialog"
)

func LocateSteamFolder() error {
	var err error

	definitions.SteamFolder, err = dialog.Directory().Title("Select steam folder").Browse()
	if err != nil {
		return err
	}
	fmt.Printf("\n Steam folder located successfully\n")

	return nil
}
