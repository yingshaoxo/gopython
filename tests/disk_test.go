package main

import (
	"strings"
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

func Test_exists(t *testing.T) {
	path := "./"
	result := disk_tool.Exists(path)
	if result == false {
		t.Fatalf("'./' should exists")
	}
}

func Test_get_absolute_path_function(t *testing.T) {
	path := "./"
	result := disk_tool.Get_absolute_path(path)

	if strings.HasPrefix(result, "/") {
		//ok
	} else {
		t.Fatalf("It should get an absolute path start with '/'")
	}

	path = "~/.auto_everything"
	result = disk_tool.Get_absolute_path(path)
	// it should be /Users/yingshaoxo/.auto_everything in one of my computers

	if (!strings.HasPrefix(result, "/")) || strings.Contains(result, "~") {
		t.Fatalf("Something is wrong")
	}
}
