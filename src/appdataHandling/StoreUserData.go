package appdataHandling

import (
	"LDT/src/appdata/userdata"
	"LDT/src/appdata/userdata/assemblyData"
	"LDT/src/functions/fileManagement"
)

func StoreUserData() error {
	err := fileManagement.StoreAsJson(fileManagement.GetPathInAppData("userdata/General.json"), userdata.General)
	if err != nil {
		return err
	}

	err = fileManagement.StoreAsJson(fileManagement.GetPathInAppData("userdata/assemblyData/InstalledAssemblies.json"), assemblyData.InstalledAssemblies)
	if err != nil {
		return err
	}

	err = fileManagement.StoreAsJson(fileManagement.GetPathInAppData("userdata/assemblyData/CurrentAssembly.json"), assemblyData.CurrentAssembly)
	if err != nil {
		return err
	}

	return nil
}
