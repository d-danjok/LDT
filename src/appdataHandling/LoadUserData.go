package appdataHandling

import (
	"LDT/src/appdata/userdata"
	"LDT/src/appdata/userdata/assemblyData"
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
		fileManagement.GetPathInAppData("userdata/assemblyData/InstalledAssemblies.json"),
		&assemblyData.InstalledAssemblies,
	)

	if err != nil && fileManagement.Exists(fileManagement.GetPathInAppData("userdata/assemblyData/InstalledAssemblies.json")) {
		return err
	}

	err = fileManagement.ReadFromJSON(
		fileManagement.GetPathInAppData("userdata/assemblyData/CurrentAssembly.json"),
		&assemblyData.CurrentAssembly,
	)

	if err != nil && fileManagement.Exists(fileManagement.GetPathInAppData("userdata/assemblyData/CurrentAssembly.json")) {
		return err
	}

	return nil
}
