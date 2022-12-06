package disk_tool

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/yingshaoxo/gopython/built_in_functions"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func Get_absolute_path(path string) string {
	usr, _ := user.Current()
	home_directory := usr.HomeDir

	var absolute_path string

	if path == "~" {
		absolute_path = home_directory
	} else if strings.HasPrefix(path, "~/") {
		absolute_path = filepath.Join(home_directory, path[2:])
	} else {
		absolute_path, _ = filepath.Abs(path)
	}

	return absolute_path
}

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
