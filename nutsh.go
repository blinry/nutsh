package main

import (
	"fmt"
	"regexp"
	"morr.cc/nutsh.git/cli"
	"morr.cc/nutsh.git/tutorial"
)

func main() {
	c := cli.Spawn("bash")
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
			fmt.Print(output)
		}

		output = c.Query("pwd")
		if ! regexp.MustCompile("/tmp/nutsh").MatchString(output) {
			tutorial.Say("stay here!")
			c.Query("cd $ROOT")
		}

		output = c.Query("test -d $ROOT/ziel && echo exists || echo nope")
		if regexp.MustCompile("nope").MatchString(output) {
			tutorial.Say("have a new dir")
			c.Query("mkdir $ROOT/ziel")
		}

		output = c.Query("test $(cat $ROOT/datei) = secret && echo ok")
		origExists := regexp.MustCompile("ok").MatchString(output)

		output = c.Query("test $(cat $ROOT/ziel/datei) = secret && echo ok")
		targetExists := regexp.MustCompile("ok").MatchString(output)

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
