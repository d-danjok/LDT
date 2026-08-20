package manifestDownload

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/sys/windows"
)

// filterAndPrintQR reads DepotDownloader stderr and prints only the QR block
func filterAndPrintQR(r io.Reader) {
	scanner := bufio.NewScanner(r)
	inQRBlock := false

	if runtime.GOOS == "windows" {
		windows.SetConsoleOutputCP(65001) // UTF-8
		windows.SetConsoleCP(65001)
	}

	for scanner.Scan() {
		line := scanner.Text()

		// DepotDownloader prints "Scan the QR code" before the block
		if strings.Contains(line, "QR") || strings.Contains(line, "qr") {
			inQRBlock = true
			fmt.Println(line) // print the instruction line too
			continue
		}

		// QR block lines contain unicode block characters or are dense ASCII
		if inQRBlock {
			// Empty line after the QR block signals it's done
			if strings.Contains(line, "Done!") {
				inQRBlock = false
				continue
			}
			fmt.Println(line)
		}
	}
}

// runDepotDownloader runs DepotDownloader with the provided arguments.
func runDepotDownloader(args ...string) error {
	// Auto-download DepotDownloader if missing
	binaryPath, err := DownloadAndExtractDepotDownloader()
	if err != nil {
		return fmt.Errorf("failed to get DepotDownloader: %w", err)
	}

	cmd := exec.Command(binaryPath, args...)

	stdout, err := cmd.StdoutPipe()

	go filterAndPrintQR(stdout)

	return cmd.Run()
}

// DownloadManifest downloads a specific depot manifest using DepotDownloader.
// App ID: 1966720, Depot ID: 1966721
func DownloadManifest(manifestID string, destDir string) error {
	return runDepotDownloader(
		"-app", "1966720",
		"-depot", "1966721",
		"-manifest", manifestID,
		"-dir", destDir,
		"-qr",
	)
}
