package main

import (
	"fmt"
	"nutsh/cli"
)

func main() {
	c := cli.Spawn(cli.BashTarget())

	for {
		cmd := c.ReadCommand()
		fmt.Printf("\n\n    you entered: %q\n\n", cmd)

		output := c.ReadOutput()
		fmt.Printf("\n\n    the output was: %q\n\n", output)
	}
}
