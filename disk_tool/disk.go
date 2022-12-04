package disk_tool

import (
	"os"
)

func Get_current_working_directory() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}
