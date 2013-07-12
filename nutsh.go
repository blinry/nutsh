package main

import "nutsh/cli"

func main() {
	c := cli.Spawn(cli.BashTarget())

	for {
		output := c.ReadOutput()
		print("output: "+output)
		cmd := c.ReadCommand()
		print("cmd: "+cmd)
	}
}
