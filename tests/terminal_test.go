package main

import (
	"testing"

	"github.com/yingshaoxo/gopython/terminal_tool"
)

func Test_run_command(t *testing.T) {
	result := terminal_tool.Run_command("uname -a")

	if len(result) == 0 {
		t.Fatalf("This function should get the shell command result back!")
	}
}
