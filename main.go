package main

import (
	definitions "LCAD/src"
	"LCAD/src/cli"
	"LCAD/src/installs"
	"LCAD/src/preloads"
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-a" {
			definitions.CleanMode = false
		}
	}
	cli.ClearTerminal()

	preloads.PreloadAssemblyList()

	installationModes := []string{"" +
		"Complete assembly installation \n " +
		"  \t(installs complete mod assembly with mods from remote source, requires steam authorisation using QR code)",
		"Individual mod installation \n" +
			"  \t(installing mods from Thunderstore by link, can be installed as new game instance or into existing installation folder)"}

	installationMode, err := cli.SelectByNum("installation type", 2, nil, installationModes)
	if err != nil {
		fmt.Printf("Error selecting installation type: %v\n", err)
		return
	}

	cli.ClearTerminal()

	switch installationMode {
	case 0:
		err = installs.InstallCompleteAssembly()
	case 1:
		err = installs.InstallWithIndividualMods()
	default:
		err = errors.New("invalid installation type")
	}
	if err != nil {
		fmt.Printf("Error assembly: %v\n", err)
		return
	}
}
