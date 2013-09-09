package main

import (
	"fmt"
	"morr.cc/nutsh.git/cli"
	"regexp"
	"strconv"
)

var (
	c              cli.CLI
	cmd, out       string
	wasInteractive bool
	didOutput      bool
)

func main() {
	c = cli.Spawn("ruby")

	/*
		say("assign 42*23 to the variable x")
		for prompt() {
			if query("x", "966") {
				say("good")
				break
			}
			output()
		}
	*/

	say("Write a method `sort` that takes an array of integers as a parameter and returns it sorted. Demonstrate it's use with an example.")
	for prompt() {
		output()
		if !command("def") && command("sort.*[\\d.*]") {
			if query("sort([2,3,1])", "NoMethodError") {
				say("The method isn't defined yet.")
			}
			if query("sort([2,3,1])", "\\[1, 2, 3\\]") {
				say("!")
				break
			}
		}
	}

	say("End of script, entering a free loop now.")
	for prompt() {
		output()
	}
}

func query(query string, expression string) bool {
	output := c.Query(query)
	return regexp.MustCompile(expression).MatchString(output)
}

func queryReturn(query string) int {
	c.Query(query)
	output := c.Query("echo $?")
	value, _ := strconv.Atoi(output[0 : len(output)-2])
	return value
}

func test(expression string) bool {
	return queryReturn("test "+expression) == 0
}

func execute(command string) {
	output := c.Query(command)
	valuestring := c.Query("echo $?")
	value, _ := strconv.Atoi(valuestring[0 : len(valuestring)-2])
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
