package cli

import (
    "os/exec"
    "strings"
)

type target struct {
    spawnCmd string
}

func stringToCmd(s string) *exec.Cmd {
    command_components := strings.Split(s, " ")
    return exec.Command(command_components[0])
}
