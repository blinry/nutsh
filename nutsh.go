package main

import (
	"fmt"
	"regexp"
	"strconv"
	"morr.cc/nutsh.git/cli"
	"morr.cc/nutsh.git/tutorial"
)

var c cli.CLI

func main() {
	c = cli.Spawn("bash")
	var (
		cmd, output string
		wasInteractive bool
	)

	c.Query("rm -rf /tmp/nutsh")
	c.Query("mkdir /tmp/nutsh")
	c.Query("cd /tmp/nutsh")
	c.Query("mkdir ziel")
	c.Query("echo secret > datei")
	c.Query("ROOT=/tmp/nutsh")

	tutorial.Say("Verschiebe `datei` in `ziel/`.")
	for {
		cmd = c.ReadCommand()
		output, wasInteractive = c.ReadOutput()
		if (! wasInteractive) {
			if regexp.MustCompile("^rm").MatchString(cmd) && !queryTrue("test -f $ROOT/datei") {
				tutorial.Say("blitz!")
				fmt.Print(output)
			} else if regexp.MustCompile("^echo").MatchString(cmd) {
				tutorial.Say("hall:")
				fmt.Print(output)
				tutorial.Say("hall aus")
			} else {
				fmt.Print(output)
			}
		}

		if regexp.MustCompile("(help|hilfe)").MatchString(cmd) {
			tutorial.Say("To move a file, use `mv`. Look at the manual to learn more.")
		}

		if !queryMatch("pwd", "/tmp/nutsh") {
			tutorial.Say("stay here!")
			c.Query("cd $ROOT")
		}

		if !queryTrue("test -d $ROOT/ziel") {
			tutorial.Say("have a new dir")
			c.Query("mkdir $ROOT/ziel")
		}

		origExists := queryTrue("test $(cat $ROOT/datei) = secret")
		targetExists := queryTrue("test $(cat $ROOT/ziel/datei) = secret")

		if targetExists {
			if origExists {
				tutorial.Say("orig still exists. remove.")
			} else {
				tutorial.Say("well done.")
				break
			}
		} else {
			if origExists {
				// nothing changed
			} else {
				tutorial.Say("have a new one.")
				c.Query("echo secret > $ROOT/datei")
			}
		}
	}

	/*
	var name string
	tutorial.Say("Please tell me your name.")
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
			tutorial.Say(fmt.Sprintf("Hello, %s!", name))
			break
		}
	}

	tutorial.Say("Now, create an alias `iam` that outputs your name.")
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
			tutorial.Say("Great.")
			break
		}
	}
	c.Query("alias iam=\"echo seb\"\n")

	tutorial.Say("Now, use this alias to pipe your name into the file `name`.")
	for {
		cmd = c.ReadCommand()
		output, wasInteractive = c.ReadOutput()
		if (! wasInteractive) {
			fmt.Print(output)
		}

		output = c.Query("test -f name && echo exists\n")
		if regexp.MustCompile("exists").MatchString(output) {
			tutorial.Say("file exists")
			output = c.Query("test $(cat name) = $(iam) && echo exists\n")
			if regexp.MustCompile("exists").MatchString(output) {
				tutorial.Say("correct content. good.")
				break
			}
		}
	}
	*/

	tutorial.Say("End of script, entering a free loop now.")
	for {
		cmd = c.ReadCommand()
		fmt.Printf("[35m\n\n    you entered: %q\n\n[0m", cmd)

		output, wasInteractive = c.ReadOutput()

		if (! wasInteractive) {
			fmt.Print("[36m    command was non-interactive\n\n[0m")
			fmt.Print(output)
		} else {
			fmt.Print("[36m\n    command was interactive\n[0m")
		}

		pwd := c.Query("echo $?\n")
		fmt.Printf("[36m\n    Command returned %s[0m", pwd)

		fmt.Printf("[32m\n    the output was: %q\n\n[0m", output)
	}
}

func queryMatch(query string, expression string) bool {
	output := c.Query(query)
	return regexp.MustCompile(expression).MatchString(output)
}

func queryReturn(query string) int {
	c.Query(query)
	output := c.Query("echo $?")
	value, _ := strconv.Atoi(output[0:len(output)-2])
	return value
}

func queryTrue(query string) bool {
	return queryReturn(query) == 0
}
