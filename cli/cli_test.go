package cli

import (
	"testing"
	"time"
)

func spawnBash() CLI {
	t := BashTarget()
	return Spawn(t)
}

func spawnRuby() CLI {
	t := RubyTarget()
	return Spawn(t)
}

func queryTest(t *testing.T, c CLI, command string, output string) {
	o := c.Query(command)
	if o != output {
		t.Fatalf("Expected %q, got %q.", output, o)
	}
}

func TestBashQueries(t *testing.T) {
	b := spawnBash()

	queryTest(t, b, "echo hi\n", "hi\r\n")
	time.Sleep(time.Millisecond)
	queryTest(t, b, "echo $((1+1))\n", "2\r\n")
}

func TestRubyQueries(t *testing.T) {
	b := spawnRuby()

	queryTest(t, b, "1+1\n", "2\r\n")
}
