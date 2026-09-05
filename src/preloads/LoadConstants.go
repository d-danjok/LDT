package preloads

import (
	definitions "LDT/src"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

func pathOf(itemName string) string {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		panic(err)
	}
	return filepath.Join(filepath.Dir(execPath), "appdata/definitions/", itemName)
}

func LoadConstants() error {
	err := readFromJSON(pathOf("LCVersions.json"), &definitions.LCVersions)
	if err != nil {
		return err
	}

	err = readFromJSON(pathOf("Assemblies.json"), &definitions.Assemblies)
	if err != nil {
		return err
	}

	unparsed, err := os.ReadFile(pathOf("general.def"))
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
