package main

import (
	. "morr.cc/nutsh.git/dsl"
	"morr.cc/nutsh.git/dsl/bash"
)

func main() {
	Spawn("bash")

	bash.Execute("rm -rf /tmp/nutsh")
	bash.Execute("mkdir /tmp/nutsh")
	bash.Execute("cd /tmp/nutsh")
	bash.Execute("ROOT=/tmp/nutsh")

	Say("Create `findgrep` here")
	for prompt() {
		if bash.Test("-f $ROOT/findgrep") {
			break
		}
	}

	Say("Make executable")
	for prompt() {
		if bash.Test("-z $($ROOT/findgrep | grep \"not executable\")") {
			break
		} else {
			Say("not executable")
		}
	}

	Say("Add Shebang, add echo test.")
	for prompt() {
		if Command("\\./findgrep") {
			if bash.Test("$(cat $ROOT/findgrep | head -1) != '#!/bin/bash'") {
				Say("Wrong shebang")
			} else {
				if bash.Test("$(./findgrep) == \"test\"") {
					Say("nice!")
				} else {
					Say("Wrong output")
				}
			}
		}
	}
}

func prompt() bool {
	Prompt()
	Output()
	if bash.Test("$(pwd) != /tmp/nutsh") {
		Say("Stay here!")
		bash.Execute("cd $ROOT")
	}
	return true
}
