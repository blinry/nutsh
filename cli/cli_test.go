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

func equalTest(t *testing.T, s1, s2 interface{}) {
	if s1 != s2 {
		t.Fatalf("Expected %q, got %q.", s2, s1)
	}
}

func queryTest(t *testing.T, c CLI, command string, output string) {
	o := c.Query(command)
	equalTest(t, o, output)
}

func TestBashQueries(t *testing.T) {
	b := spawnBash()

	queryTest(t, b, "echo hi\n", "hi\r\n")
	time.Sleep(time.Millisecond)
	queryTest(t, b, "echo $((1+1))\n", "2\r\n")
}

func TestBashTabCompletion(t *testing.T) {
	b := spawnBash()

	queryTest(t, b, "echo /t\t\n", "/tmp/\r\n")
}

func TestBashEditing(t *testing.T) {
	b := spawnBash()
	queryTest(t, b, "echo notthisfuhi\n", "hi\r\n")
}

func TestBashHistory(t *testing.T) {
	b := spawnBash()
	queryTest(t, b, "echo rememberme\n", "rememberme\r\n")
	time.Sleep(time.Millisecond)
	queryTest(t, b, "OA\n", "rememberme\r\n")
}

func TestBashLoop(t *testing.T) {
	c := spawnBash()

	c.read(promptType)
	c.send("echo flupp\n")
	cmd := c.read(commandType)
	equalTest(t, cmd, "echo flupp")
	output := c.read(outputType)
	equalTest(t, output, "flupp\r\n")
}

func TestBashEmpty(t *testing.T) {
	c := spawnBash()
	queryTest(t, c, "\n", "")
}

func TestRubyQueries(t *testing.T) {
	b := spawnRuby()

	queryTest(t, b, "1+1\n", "2\r\n")
}
