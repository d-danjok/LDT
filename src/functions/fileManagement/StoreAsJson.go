package fileManagement

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func StoreAsJson[t interface{}](filePath string, data t) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if !Exists(filePath) {
		err = os.MkdirAll(filepath.Dir(filePath), 0644)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(filePath, jsonData, 0644)
}
