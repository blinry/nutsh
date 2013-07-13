package cli

import (
	"os/exec"
	"strings"
)

type target struct {
	spawnCmd string
	initCmd string
	markedPrompt bool
	prompt string
}

func stringToCmd(s string) *exec.Cmd {
	command_components := strings.Split(s, " ")
	if len(command_components) == 1 {
		return exec.Command(command_components[0])
	} else {
        return exec.Command(command_components[0], command_components[1:len(command_components)-1]...)
	}
}

func BashTarget() target {
	return target{
		spawnCmd: "bash --norc -i",
		initCmd: `export PS1="$(echo -e "\u2603")\w $ $(echo -e "\u2603")"
`,
		markedPrompt: true,
	}
}

func RubyTarget() target {
	return target{
		spawnCmd: "irb",
		initCmd: `conf.prompt_i = "\u2603>> \u2603";0
conf.return_format = "%s\n"
`,
		//markedPrompt: false,
		//prompt: ">> ",
	}
}
