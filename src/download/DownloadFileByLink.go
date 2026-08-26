package downloads

import (
	"io"
	"net/http"
	"os"
)

func DownloadFileByLink(link string, fileName string) error {
	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	_ = size
	//fmt.Printf("Downloaded a file %s with size %d\n\n", fileName, size)

	return nil
}
