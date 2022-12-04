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

func Test_path_join(t *testing.T) {
	path1 := "/home"
	path2 := "yingshaoxo"
	result := disk_tool.Path_join(path1, path2)
	if result != "/home/yingshaoxo" {
		t.Fatalf("result should be '/home/yingshaoxo'")
	}
}
