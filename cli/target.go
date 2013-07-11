package cli

import (
	"os/exec"
	"strings"
)

type target struct {
	spawnCmd string
	initCmd string
}

func stringToCmd(s string) *exec.Cmd {
	command_components := strings.Split(s, " ")
	if len(command_components) == 1 {
		return exec.Command(command_components[0])
	} else {
        return exec.Command(command_components[0], command_components[1:len(command_components)-1]...)
	}
}
