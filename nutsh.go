package main

import (
	"fmt"
	"morr.cc/nutsh.git/cli"
)

func main() {
	c := cli.Spawn("bash")

	for {
		cmd := c.ReadCommand()
		fmt.Printf("[35m\n\n    you entered: %q\n\n[0m", cmd)

		output := c.ReadOutput()
		fmt.Print(output)

		fmt.Printf("[32m\n    the output was: %q\n\n[0m", output)
	}
}
