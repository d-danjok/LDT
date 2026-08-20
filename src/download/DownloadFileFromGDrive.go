package downloads

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
	"strings"
)

func extractFileID(link string) (string, error) {
	re := regexp.MustCompile(`(?:/d/|id=)([a-zA-Z0-9_-]{10,})`)
	matches := re.FindStringSubmatch(link)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid Google Drive link")
	}
	return matches[1], nil
}

func DownloadFileFromGDrive(link, destPath string) error {
	fileID, err := extractFileID(link)
	if err != nil {
		return err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("failed to create cookie jar: %w", err)
	}
	client := &http.Client{Jar: jar}

	// Modern usercontent domain with confirm=t baked in
	downloadURL := fmt.Sprintf(
		"https://drive.usercontent.google.com/download?id=%s&export=download&authuser=0&confirm=t",
		fileID,
	)

	resp, err := client.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to request file: %w", err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "text/html") {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}
		return fmt.Errorf("got HTML instead of file (first 500 chars):\n%s", string(body[:min(500, len(body))]))
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// extractConfirmURL pulls the confirmed download URL out of the warning HTML page
func extractConfirmURL(html string, fileID string) (string, error) {
	// Google embeds a confirm token like: confirm=t&uuid=...
	// Modern Drive pages use a form action or a direct href
	re := regexp.MustCompile(`confirm=([0-9A-Za-z_\-]+)`)
	matches := re.FindStringSubmatch(html)

	if len(matches) < 2 {
		// Newer Drive pages may use a UUID-based URL instead
		reUUID := regexp.MustCompile(`(/uc\?export=download[^"']+)`)
		uMatches := reUUID.FindStringSubmatch(html)
		if len(uMatches) < 2 {
			return "", fmt.Errorf("could not find confirmation token in page")
		}
		// Unescape HTML entities like &amp;
		rawURL := strings.ReplaceAll(uMatches[1], "&amp;", "&")
		return "https://drive.google.com" + rawURL, nil
	}

	confirmToken := matches[1]
	return fmt.Sprintf(
		"https://drive.google.com/uc?export=download&confirm=%s&id=%s",
		confirmToken, fileID,
	), nil
}
