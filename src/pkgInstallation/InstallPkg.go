package pkgInstallation

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var BepInExInstalled bool

func archiveName(pkgName string) string {
	return pkgName + ".zip"
}

func installBepInEx(archiveLocation string, installationPath string) error {
	const pkgName = "BepInEx-BepInExPack"

	//no more install
	if BepInExInstalled {
		return nil
	}

	reader, err := zip.OpenReader(filepath.Join(archiveLocation, archiveName(pkgName)))
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, f := range reader.File {
		//install only contents of BepInExPack subfolder
		if strings.Split(filepath.ToSlash(f.Name), "/")[0] != "BepInExPack" {
			continue
		}
		path := filepath.Join(installationPath, filepath.Join(strings.Split(filepath.ToSlash(f.Name), "/")[1:]...))

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
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

	BepInExInstalled = true
	return nil
}

func InstallPkg(pkgName string, archiveLocation string, installationPath string) error {
	if pkgName == "BepInEx-BepInExPack" {
		return installBepInEx(archiveLocation, installationPath)
	}

	reader, err := zip.OpenReader(filepath.Join(archiveLocation, archiveName(pkgName)))
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, f := range reader.File {
		//destination folder switch
		var path string
		switch strings.Split(filepath.ToSlash(f.Name), "/")[0] {
		case "BepInEx":
			path = filepath.Join(installationPath, f.Name)

		case "config", "patchers", "plugins":
			path = filepath.Join(installationPath, "BepInEx", f.Name)

		default:
			path = filepath.Join(installationPath, "BepInEx/plugins", pkgName, f.Name)
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
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
