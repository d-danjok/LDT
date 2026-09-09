package installs

import (
	"LDT/src/cli"
	"LDT/src/functions/fileManagement"
	"fmt"
)

func InstallCompleteAssembly() error {
	assemblyToInstall, err := cli.SelectAssemblyToInstall()
	if err != nil {
		return fmt.Errorf("error selecting assembly to install: %v", err)
	}

	cli.ClearTerminal()

	fmt.Printf("Installing %s", assemblyToInstall.Name)

	err = assemblyToInstall.Install(fileManagement.GetPathInAppData("assemblies"))
	if err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}

	return nil
}
