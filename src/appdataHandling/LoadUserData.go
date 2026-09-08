package appdataHandling

import (
	"LDT/src/appdata/userdata"
	"LDT/src/functions/fileManagement"
)

func LoadUserData() error {
	err := fileManagement.ReadFromJSON(
		fileManagement.GetPathInAppData("userdata/General.json"),
		&userdata.General,
	)

	if err != nil && fileManagement.Exists(fileManagement.GetPathInAppData("userdata/General.json")) {
		return err
	}

	err = fileManagement.ReadFromJSON(
		fileManagement.GetPathInAppData("userdata/InstalledAssemblies.json"),
		&userdata.InstalledAssemblies,
	)

	if err != nil && fileManagement.Exists(fileManagement.GetPathInAppData("userdata/InstalledAssemblies.json")) {
		return err
	}
	return nil
}
