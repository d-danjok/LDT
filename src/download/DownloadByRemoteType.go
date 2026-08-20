package downloads

import (
	"path/filepath"
)

func DownloadByRemoteType(remoteType string, remoteLink string, destPath string, fileName string) error {
	switch remoteType {
	case "":
		return nil
	case "GoogleDrive":
		err := DownloadFileFromGDrive(remoteLink, filepath.Join(destPath, fileName))
		if err != nil {
			return err
		}
	}
	return nil
}
