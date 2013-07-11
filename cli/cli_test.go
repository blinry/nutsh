package cli

import "testing"

func spawnBash() CLI {
	t := target{"bash --norc -i"}
	return Spawn(t)
}

func queryTest(t *testing.T, c CLI, command string, output string) {
	o := c.Query(command)
	if o != output {
		t.Fatalf("Expected %q, got %q.", output, o)
	}
}

func TestQuery(t *testing.T) {
	b := spawnBash()

	queryTest(t, b, "echo hi", "hi\r")
	queryTest(t, b, "echo $((1+1))", "2\r")
}
