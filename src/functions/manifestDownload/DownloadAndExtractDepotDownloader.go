package manifestDownload

import (
	"LDT/src/functions/extraction"
	"LDT/src/functions/fileManagement"
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
	binaryPath := filepath.Join(fileManagement.GetPathInAppData("depotDownloader"), getDepotDownloaderBinaryName())

	// Already exists, skip download
	if fileManagement.Exists(binaryPath) {
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

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		out.Close()
		return "", fmt.Errorf("failed to save zip: %w", err)
	}
	out.Close()

	// Extract zip
	err = extraction.Unzip(zipPath, filepath.Dir(binaryPath))
	if err != nil {
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
