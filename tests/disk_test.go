package main

import (
	"testing"

	"github.com/yingshaoxo/gopython/disk_tool"
)

func Test_get_pwd(t *testing.T) {
	path := disk_tool.Get_current_working_directory()
	if len(path) == 0 {
		t.Fatalf("This function should return the current path as a string!")
	}
}
