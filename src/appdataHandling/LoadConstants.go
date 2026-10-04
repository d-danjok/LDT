package appdataHandling

import (
	definitions "LDT/src/appdata"
	"LDT/src/functions/fileManagement"
	"time"
)

func LoadConstants() error {
	err := fileManagement.ReadFromJSON(fileManagement.GetPathInAppData("definitions/LCVersions.json"), &definitions.LCVersions)
	if err != nil {
		return err
	}

	for i, lcVersion := range definitions.LCVersions {
		if lcVersion.LastDate == "none" {
			definitions.LCVersions[i].LastDate = time.Now().Format("2006-01-02")
		}
	}

	return nil
}
