package fileManagement

import "path/filepath"

func GetPathInAppData(itemPath string) string {
	//execPath, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//execPath, err = filepath.EvalSymlinks(execPath)
	//if err != nil {
	//	panic(err)
	//}
	//return filepath.Join(filepath.Dir(execPath), "appdata", itemPath)
	return filepath.Join("appdata", itemPath)
}
