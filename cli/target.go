package cli

import (
	"os/exec"
	"strings"
)

type target struct {
	spawnCmd string
	initCmd  string
}

var targets = map[string]target{
	"bash": target{
		spawnCmd: "bash --norc -i",
		initCmd: ` export HISTCONTROL=ignorespace
 history -d $((HISTCMD-1))
 alias ls="ls --color=auto"
 export PS2="\[$(echo -e "\xe2\x98\x85")\]> \[\e[0m$(echo -e "\xe2\x98\x83")\]"
 export PS1="\[$(echo -e "\xe2\x98\x83")\e[34m\e[1m\]\w $ \[\e[0m$(echo -e "\xe2\x98\x83")\]"
 shopt -s checkwinsize
`,
	},
	"ruby": target{
		spawnCmd: "irb",
		initCmd: `conf.return_format = "%s\n"
conf.prompt_i = "\u2603>> \u2603";0
`,
	},
	"python": target{
		spawnCmd: "python",
		initCmd: `import sys
sys.ps2 = "\u2605... \u2603"
sys.ps1 = "\u2603>>> \u2603"
`,
	},
}

func stringToCmd(s string) *exec.Cmd {
	command_components := strings.Split(s, " ")
	if len(command_components) == 1 {
		return exec.Command(command_components[0])
	} else {
		return exec.Command(command_components[0], command_components[1:len(command_components)-1]...)
	}
}
