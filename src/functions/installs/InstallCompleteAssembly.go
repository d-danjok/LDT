package installs

import (
	definitions "LDT/src/appdata"
	"LDT/src/appdata/userdata"
	"LDT/src/cli"
	"fmt"
	"path/filepath"
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

	err = assemblyToInstall.Install(filepath.Join(userdata.General.SteamFolderLocation, definitions.LCAssembliesDefFolderSubPath))
	if err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}

	return nil
}
