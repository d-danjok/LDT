package appdataHandling

import (
	definitions "LDT/src/appdata"
	"LDT/src/appdata/userdata"
	"LDT/src/appdata/userdata/assemblyData"
	"LDT/src/functions/fileManagement"
)

func LoadUserData() error {

	if !fileManagement.Exists(fileManagement.GetPathInAppData("userdata/General.json")) ||
		!fileManagement.Exists(fileManagement.GetPathInAppData("userdata/assemblyData/InstalledAssemblies.json")) ||
		!fileManagement.Exists(fileManagement.GetPathInAppData("userdata/assemblyData/CurrentAssembly.json")) {

		definitions.IsFirstLaunch = true
		return nil
	}

	err := fileManagement.ReadFromJSON(
		fileManagement.GetPathInAppData("userdata/General.json"),
		&userdata.General,
	)

	if err != nil {
		return err
	}

	err = fileManagement.ReadFromJSON(
		fileManagement.GetPathInAppData("userdata/assemblyData/InstalledAssemblies.json"),
		&assemblyData.InstalledAssemblies,
	)

	if err != nil {
		return err
	}

	err = fileManagement.ReadFromJSON(
		fileManagement.GetPathInAppData("userdata/assemblyData/CurrentAssembly.json"),
		&assemblyData.CurrentAssembly,
	)

	if err != nil {
		return err
	}

	return nil
}
