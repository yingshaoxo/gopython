package disk_tool

import (
	"os"
	"path/filepath"
)

func Get_current_working_directory() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

func Path_join(path1 string, path2 string) string {
	return filepath.Join(path1, path2)
}
