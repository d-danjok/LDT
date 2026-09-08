package fileManagement

import (
	"encoding/json"
	"os"
)

func ReadFromJSON(filePath string, dest any) error {
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
