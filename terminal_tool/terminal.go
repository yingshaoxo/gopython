package terminal_tool

import (
	"bytes"
	"os/exec"
	"strings"
)

func Run_command(command string) string {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)

	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Run()

	return out.String()
}
