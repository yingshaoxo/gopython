package disk_tool

import (
	"os"
	"path/filepath"

	"github.com/yingshaoxo/gopython/built_in_functions"
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

func Is_directory(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		built_in_functions.Print(err.Error())
	}

	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
}

func Remove_a_file_or_folder(path string) bool {
	var err error

	if Is_directory(path) {
		err = os.RemoveAll(path)
	} else {
		err = os.Remove(path)
	}

	if err != nil {
		built_in_functions.Print(err.Error())
		return false
	} else {
		return true
	}
}
