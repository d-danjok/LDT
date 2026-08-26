package installs

import (
	definitions "LDT/src"
	"LDT/src/cli"
	"fmt"
)

func InstallCompleteAssembly() error {
	assemblyToInstall, err := cli.SelectAssemblyToInstall()
	if err != nil {
		return fmt.Errorf("error selecting assembly to install: %v", err)
	}

	err = cli.LocateSteamFolder()
	if err != nil {
		return fmt.Errorf("error locating Steam folder: %v", err)
	}

	fmt.Printf("Installing %s", assemblyToInstall.Name)

	err = assemblyToInstall.Install(definitions.SteamFolder)
	if err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}

	return nil
}
