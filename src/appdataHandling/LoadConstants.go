package appdataHandling

import (
	definitions "LDT/src/appdata"
	"LDT/src/functions/fileManagement"
	"fmt"
	"os"
	"strings"
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

	unparsed, err := os.ReadFile(fileManagement.GetPathInAppData("definitions/general.def"))
	if err != nil {
		return err
	}
	//need to split twice because windows standard on creating new line is \r\n, while UNIX-based systems use \n
	fileLines := strings.Split(string(unparsed), "\r\n")
	if len(fileLines) < 2 {
		fileLines = strings.Split(string(unparsed), "\n")
	}

	for _, line := range fileLines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " : ")
		switch parts[0] {
		case "SteamFolder":
			definitions.SteamFolder = parts[1]
		case "LCAssembliesDefFolderSubPath":
			definitions.LCAssembliesDefFolderSubPath = parts[1]
		default:
			return fmt.Errorf("unrecognized constant: %s", parts[0])
		}
	}

	return nil
}
