package fsManagement

import (
	"os"
)

func Exists(locationPath string) bool {
	_, err := os.Stat(locationPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	// For other errors treat as existing but inaccessible
	return true
}
