package extraction

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(archivePath string, destPath string) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}

	// ensure tmp files will be closed and deleted
	defer func(name string) {
		err := r.Close()
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		err = os.Remove(name)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}(archivePath)

	for _, f := range r.File {
		path := filepath.Join(destPath, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.Create(path)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		srcFile, err := f.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		if _, err = io.Copy(dstFile, srcFile); err != nil {
			return err
		}
	}

	return nil
}
