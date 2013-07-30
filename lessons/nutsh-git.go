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
	bash.Execute("ROOT=$(pwd)")

	decide:
	Say("Yes or no?")
	for Prompt() {
		Output()
		if Command("[yY]es") {
			goto yes
		}
		if Command("[nN]o") {
			goto no
		}
	}

	yes:
	Say("You sure?")
	for Prompt() {
		Output()
		if Command("[yY]es") {
			Say("OK")
			return
		}
		if Command("[nN]o") {
			goto decide
		}
	}

	no:
	Say("Or rather yes?")
	for Prompt() {
		Output()
		if Command("[yY]es") {
			goto yes
		}
		if Command("[nN]o") {
			Say("OK, no.")
			return
		}
	}

	Say("Hi! Welcome to the Git crash course Initialize a new repository with `git init`.")
	for Prompt() {
		Output()
		if bash.Test("-f $ROOT/.git/HEAD") {
			break
		}
	}

	Say("Now, let's have a look at what happened. `git init` created a `.git/` directory where all stuff related to Git lives. Please go there.")
	for Prompt() {
		Output()
		if bash.Test("$(pwd) = \"$ROOT/.git\"") {
			break
		}
	}

	Say("(explanation, moving back)")
	bash.Execute("cd ..")


	Say("Add a file, do a commit")
	for Prompt() {
		Output()
		if QueryOutput("git log", "Author:") {
			Say("nice.")
			break
		}
	}

	Say("End of script, entering a free loop now.")
	for Prompt() {
		Output()
	}
}
