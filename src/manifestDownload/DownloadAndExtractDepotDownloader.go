package manifestDownload

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func getDepotDownloaderURL() string {
	base := "https://github.com/SteamRE/DepotDownloader/releases/latest/download/"
	switch runtime.GOOS {
	case "windows":
		return base + "DepotDownloader-windows-x64.zip"
	case "darwin":
		return base + "DepotDownloader-macos-x64.zip"
	default:
		return base + "DepotDownloader-linux-x64.zip"
	}
}

func getDepotDownloaderBinaryName() string {
	if runtime.GOOS == "windows" {
		return "DepotDownloader.exe"
	}
	return "DepotDownloader"
}

// DownloadAndExtractDepotDownloader downloads the binary into Temp folder
func DownloadAndExtractDepotDownloader() (string, error) {
	binaryPath := filepath.Join(os.TempDir(), getDepotDownloaderBinaryName())

	// Already exists, skip download
	if _, err := os.Stat(binaryPath); err == nil {
		return binaryPath, nil
	}

	url := getDepotDownloaderURL()
	fmt.Printf("Downloading DepotDownloader from %s...\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	zipPath := filepath.Join(os.TempDir(), "depotdownloader.zip")
	out, err := os.Create(zipPath)
	if err != nil {
		return "", fmt.Errorf("failed to create zip: %w", err)
	}

	if _, err = io.Copy(out, resp.Body); err != nil {
		out.Close()
		return "", fmt.Errorf("failed to save zip: %w", err)
	}
	out.Close()

	// Extract zip
	if err := unzip(zipPath, os.TempDir()); err != nil {
		return "", fmt.Errorf("failed to extract: %w", err)
	}
	os.Remove(zipPath)

	// Make executable on Linux/macOS
	if runtime.GOOS != "windows" {
		os.Chmod(binaryPath, 0755)
	}

	fmt.Println("DepotDownloader ready at:", binaryPath)
	return binaryPath, nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}

		out, err := os.Create(path)
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			out.Close()
			return err
		}

		_, err = io.Copy(out, rc)
		rc.Close()
		out.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
