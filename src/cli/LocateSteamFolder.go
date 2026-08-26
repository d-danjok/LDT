package cli

import (
	definitions "LDT/src"
	"LDT/src/gui"
	"fmt"
	"runtime"
)

func LocateSteamFolder() error {
	var steamFolderLocated bool
	var err error

	if runtime.GOOS == "windows" {
		definitions.SteamFolder = "C:\\Program Files (x86)\\Steam"
		steamFolderLocated, err = Confirm("use default Steam folder location (" + definitions.SteamFolder + ")")
		if err != nil {
			return err
		}
		fmt.Printf("\n")
	}
	if !steamFolderLocated {
		if runtime.GOOS == "windows" {
			fmt.Printf("\nMoving to manual selection\n")
		}
		err = gui.BrowseForSteamFolder()
		if err != nil {
			return err
		}
	}

	return nil
}
