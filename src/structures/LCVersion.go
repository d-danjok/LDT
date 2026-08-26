package structures

import (
	"LCAD/src/manifestDownload"
	"fmt"
	"os"
)

type LCVersion struct {
	Name       string
	ManifestID string
}

func (v LCVersion) Install(installationPath string) error {
	err := os.MkdirAll(os.TempDir(), os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Printf("\n")

	//download base
	err = manifestDownload.DownloadManifest(
		v.ManifestID,
		installationPath,
	)

	if err != nil {
		return err
	}

	return nil
}
