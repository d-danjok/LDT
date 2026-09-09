package appdataHandling

import (
	definitions "LDT/src/appdata"
	"LDT/src/functions/fileManagement"
)

func LoadConstants() error {
	err := fileManagement.ReadFromJSON(fileManagement.GetPathInAppData("definitions/LCVersions.json"), &definitions.LCVersions)
	if err != nil {
		return err
	}

	err = fileManagement.ReadFromJSON(fileManagement.GetPathInAppData("definitions/Assemblies.json"), &definitions.Assemblies)
	if err != nil {
		return err
	}

	return nil
}
