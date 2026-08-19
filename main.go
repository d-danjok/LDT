package main

import (
	"LCAD/src"
	"LCAD/src/cli"
	"LCAD/src/gui"
	"LCAD/src/preloads"
	"fmt"
	"runtime"
)

func main() {
	var steamFolderLocated bool

	preloads.PreloadAssemblyList()

	assemblies := definitions.Assemblies
	_ = assemblies

	if runtime.GOOS == "windows" {
		definitions.SteamFolder = "C:\\Program Files (x86)\\Steam"

		fmt.Printf("\nDo you want to use default Steam folder location (%s) [y/n] \n: ", definitions.SteamFolder)
		var confirmation string

		fmt.Scanln(&confirmation)
		if confirmation == "y" || confirmation == "Y" {
			steamFolderLocated = true
		}
		fmt.Printf("\n")
	}
	if !steamFolderLocated {
		err := gui.LocateSteamFolder()
		if err != nil {
			fmt.Printf("Error selecting folder: %v\n", err)
			return
		}
	}
	assemblyToInstall, err := cli.SelectAssemblyToInstall()
	if err != nil {
		fmt.Printf("Error selecting assembly: %v\n", err)
		return
	}

	err = assemblyToInstall.Install(definitions.SteamFolder)
	if err != nil {
		fmt.Printf("Error installing assembly: %v\n", err)
		return
	}
}
