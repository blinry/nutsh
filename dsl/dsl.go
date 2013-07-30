package dsl

import (
	"fmt"
	"regexp"
	"morr.cc/nutsh.git/cli"
)

var (
	cmdline cli.CLI
	lastCommand, lastOutput string
	wasInteractive bool
	didOutput bool
)

func Spawn(target string) {
	cmdline = cli.Spawn(target)
}

func Query(query string) string {
	return cmdline.Query(query)
}

func QueryOutput(query string, expression string) bool {
	output := cmdline.Query(query)
	return regexp.MustCompile(expression).MatchString(output)
}

func Say(text string) {
	text = regexp.MustCompile("`([^`]+)`").ReplaceAllString(text, "[32m$1[36m")
	fmt.Printf("[36m\n    %s\n\n[0m", text)
}

func Command(expression string) bool {
	return regexp.MustCompile(expression).MatchString(lastCommand)
}

func OutputMatch(expression string) bool {
	return regexp.MustCompile(expression).MatchString(lastOutput)
}

func Output() {
	if !wasInteractive && !didOutput {
		fmt.Print(lastOutput)
		didOutput = true
	}
}

func Prompt() bool {
	didOutput = false
	lastCommand = cmdline.ReadCommand()
	lastOutput, wasInteractive = cmdline.ReadOutput()
	return true
}
