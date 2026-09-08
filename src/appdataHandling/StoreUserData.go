package appdataHandling

import (
	"LDT/src/appdata/userdata"
	"LDT/src/functions/fileManagement"
)

func StoreUserData() error {
	err := fileManagement.StoreAsJson(fileManagement.GetPathInAppData("userdata/General.json"), userdata.General)
	if err != nil {
		return err
	}

	err = fileManagement.StoreAsJson(fileManagement.GetPathInAppData("userdata/InstalledAssemblies.json"), userdata.InstalledAssemblies)
	if err != nil {
		return err
	}

	return nil
}
