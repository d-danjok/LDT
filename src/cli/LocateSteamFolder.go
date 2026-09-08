package cli

import (
	definitions "LDT/src/appdata"
	"LDT/src/appdata/userdata"
	"LDT/src/gui"
	"fmt"
	"runtime"
)

func LocateSteamFolder() error {
	var steamFolderLocated bool
	var err error

	if userdata.General.SteamFolderLocation != "" {
		steamFolderLocated, err = Confirm("use previously selected Steam folder location (" + userdata.General.SteamFolderLocation + ")")
		if err != nil {
			return err
		}
		fmt.Printf("\n")
	}

	if steamFolderLocated {
		return nil
	}

	if runtime.GOOS == "windows" {
		steamFolderLocated, err = Confirm("use default Steam folder location (" + definitions.SteamFolder + ")")
		if err != nil {
			return err
		}
		fmt.Printf("\n")
		if steamFolderLocated {
			userdata.General.SteamFolderLocation = definitions.SteamFolder
			return nil
		}
	}

	if runtime.GOOS == "windows" {
		fmt.Printf("\nMoving to manual selection\n")
	}
	err = gui.BrowseForSteamFolder()
	if err != nil {
		return err
	}

	return nil
}
