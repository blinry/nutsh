package main

import (
	"fmt"
	"regexp"
	"strconv"
	"morr.cc/nutsh.git/cli"
)

var (
	c cli.CLI
	cmd, out string
	wasInteractive bool
	didOutput bool
)

func main() {
	c = cli.Spawn("bash")

	execute("rm -rf /tmp/nutsh")
	execute("mkdir /tmp/nutsh")
	execute("cd /tmp/nutsh")
	execute("mkdir ziel")
	execute("echo secret > datei")
	execute("ROOT=/tmp/nutsh")

	say("Verschiebe `datei` in `ziel/`.")
	for prompt() {
		if command("^echo") {
			say("hall:")
			output()
			say("hall aus")
		}

		output()

		if command("(help|hilfe)") {
			say("To move a file, use `mv`. Look at the manual to learn more.")
		}

		if !query("pwd", "/tmp/nutsh") {
			say("stay here!")
			execute("cd $ROOT")
		}

		if !test("-d $ROOT/ziel") {
			say("have a new dir")
			execute("mkdir $ROOT/ziel")
		}

		origExists := test("$(cat $ROOT/datei) = secret")
		targetExists := test("$(cat $ROOT/ziel/datei) = secret")

		if targetExists {
			if origExists {
				say("orig still exists. remove.")
			} else {
				say("well done.")
				break
			}
		} else {
			if origExists {
				// nothing changed
			} else {
				say("have a new one.")
				execute("echo secret > $ROOT/datei")
			}
		}
	}

	/*
	var name string
	say("Please tell me your name.")
	for {
		cmd = c.ReadCommand()
		output, wasInteractive = c.ReadOutput()
		if (! wasInteractive) {
			fmt.Print(output)
		}

		iam := regexp.MustCompile("i am ([\\w]+)")
		m := iam.FindStringSubmatch(output)
		if m != nil {
			name = m[1]
			say(fmt.Sprintf("Hello, %s!", name))
			break
		}
	}

	say("Now, create an alias `iam` that outputs your name.")
	for {
		cmd = c.ReadCommand()
		output, wasInteractive = c.ReadOutput()
		if (! wasInteractive) {
			fmt.Print(output)
		}

		output = c.Query("iam\n")
		iam := regexp.MustCompile(name)
		m := iam.FindStringSubmatch(output)
		if m != nil {
			say("Great.")
			break
		}
	}
	c.Query("alias iam=\"echo seb\"\n")

	say("Now, use this alias to pipe your name into the file `name`.")
	for {
		cmd = c.ReadCommand()
		output, wasInteractive = c.ReadOutput()
		if (! wasInteractive) {
			fmt.Print(output)
		}

		output = c.Query("test -f name && echo exists\n")
		if regexp.MustCompile("exists").MatchString(output) {
			say("file exists")
			output = c.Query("test $(cat name) = $(iam) && echo exists\n")
			if regexp.MustCompile("exists").MatchString(output) {
				say("correct content. good.")
				break
			}
		}
	}
	*/

	say("End of script, entering a free loop now.")
	for {
		cmd = c.ReadCommand()
		fmt.Printf("[35m\n\n    you entered: %q\n\n[0m", cmd)

		out, wasInteractive = c.ReadOutput()

		if (! wasInteractive) {
			fmt.Print("[36m    command was non-interactive\n\n[0m")
			fmt.Print(out)
		} else {
			fmt.Print("[36m\n    command was interactive\n[0m")
		}

		pwd := c.Query("echo $?\n")
		fmt.Printf("[36m\n    Command returned %s[0m", pwd)

		fmt.Printf("[32m\n    the output was: %q\n\n[0m", out)
	}
}

func query(query string, expression string) bool {
	output := c.Query(query)
	return regexp.MustCompile(expression).MatchString(output)
}

func queryReturn(query string) int {
	c.Query(query)
	output := c.Query("echo $?")
	value, _ := strconv.Atoi(output[0:len(output)-2])
	return value
}

func test(expression string) bool {
	return queryReturn("test "+expression) == 0
}

func execute(command string) {
	output := c.Query(command)
	valuestring := c.Query("echo $?")
	value, _ := strconv.Atoi(valuestring[0:len(valuestring)-2])
	if value != 0 {
		panic(fmt.Sprintf("executing `%s` failed: %s", command, output))
	}
}

func say(text string) {
	text = regexp.MustCompile("`([^`]+)`").ReplaceAllString(text, "[32m$1[36m")
	fmt.Printf("[36m\n\n    %s\n\n[0m", text)
}

func command(expression string) bool {
	return regexp.MustCompile(expression).MatchString(cmd)
}

func output() {
	if !wasInteractive && !didOutput {
		fmt.Print(out)
		didOutput = true
	}
}

func prompt() bool {
	didOutput = false
	cmd = c.ReadCommand()
	out, wasInteractive = c.ReadOutput()
	return true
}
