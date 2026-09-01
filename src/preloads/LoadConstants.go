package preloads

import (
	definitions "LDT/src"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func readFromJSON(filePath string, dest any) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &dest)
	if err != nil {
		return err
	}

	return nil
}

func LoadConstants() error {
	err := readFromJSON("appdata/definitions/LCVersions.json", &definitions.LCVersions)
	if err != nil {
		return err
	}

	err = readFromJSON("appdata/definitions/Assemblies.json", &definitions.Assemblies)
	if err != nil {
		return err
	}

	unparced, err := os.ReadFile("appdata/definitions/general.def")
	if err != nil {
		return err
	}
	//need to split twice because windows standard on creating new line is \r\n, while UNIX-based systems use \n
	fileLines := strings.Split(string(unparced), "\r\n")
	if len(fileLines) < 2 {
		fileLines = strings.Split(string(unparced), "\n")
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
