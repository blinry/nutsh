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
	bash.Execute("mkdir ziel")
	bash.Execute("echo secret > datei")
	bash.Execute("ROOT=/tmp/nutsh")

	Say("Verschiebe `datei` in `ziel/`.")
	for Prompt() {
		if Command("^echo") {
			Say("hall:")
			Output()
			Say("hall aus")
		}

		Output()

		if Command("(help|hilfe)") {
			Say("To move a file, use `mv`. Look at the manual to learn more.")
		}

		if !QueryOutput("pwd", "/tmp/nutsh") {
			Say("stay here!")
			bash.Execute("cd $ROOT")
		}

		if !bash.Test("-d $ROOT/ziel") {
			Say("have a new dir")
			bash.Execute("mkdir $ROOT/ziel")
		}

		origExists := bash.Test("$(cat $ROOT/datei) = secret")
		targetExists := bash.Test("$(cat $ROOT/ziel/datei) = secret")

		if targetExists {
			if origExists {
				Say("orig still exists. remove.")
			} else {
				Say("well done.")
				break
			}
		} else {
			if origExists {
				// nothing changed
			} else {
				Say("have a new one.")
				bash.Execute("echo secret > $ROOT/datei")
			}
		}
	}

	Say("Please tell me your name.")
	for Prompt() {
		Output()

		if OutputMatch("i am (\\w+)") {
			Say("Hello, $0!")
			bash.Execute("NAME=$0")
			break
		}
	}

	/*
		Say("Now, create an alias `iam` that outputs your name.")
		for Prompt() {
			output = c.Query("iam\n")
			iam := regexp.MustCompile(name)
			m := iam.FindStringSubmatch(output)
			if m != nil {
				Say("Great.")
				break
			}
		}
		c.Query("alias iam=\"echo seb\"\n")

		Say("Now, use this alias to pipe your name into the file `name`.")
		for {
			cmd = c.ReadCommand()
			output, wasInteractive = c.ReadOutput()
			if (! wasInteractive) {
				fmt.Print(output)
			}

			output = c.Query("test -f name && echo exists\n")
			if regexp.MustCompile("exists").MatchString(output) {
				Say("file exists")
				output = c.Query("test $(cat name) = $(iam) && echo exists\n")
				if regexp.MustCompile("exists").MatchString(output) {
					Say("correct content. good.")
					break
				}
			}
		}
	*/

	Say("End of script, entering a free loop now.")
	for Prompt() {
		Output()
	}
}
