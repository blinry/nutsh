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
		initCmd: `export PS1="$(echo -e "\u2603")\e[34m\e[1m\w $ \e[0m$(echo -e "\u2603")"
`,
		markedPrompt: true,
	}
}

func RubyTarget() target {
	return target{
		spawnCmd: "irb",
		initCmd: `conf.return_format = "%s\n"
conf.prompt_i = "\u2603>> \u2603";0
`,
		//markedPrompt: false,
		//prompt: ">> ",
	}
}
