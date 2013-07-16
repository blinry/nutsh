package cli

import (
	"testing"
	//"time"
)

func spawnBash() CLI {
	return Spawn("bash")
}

func spawnRuby() CLI {
	return Spawn("ruby")
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
	queryTest(t, b, "echo $((1+1))\n", "2\r\n")
}

func TestBashStressQueries(t *testing.T) {
	b := spawnBash()

	for i := 0; i<100; i++ {
		queryTest(t, b, "echo hi\n", "hi\r\n")
	}
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
	queryTest(t, b, "OA\n", "rememberme\r\n")
}

func TestBashLoop(t *testing.T) {
	c := spawnBash()
	c.allowInteractivity = false

	c.read(promptType)
	c.send("echo flupp\n")
	cmd := c.read(finalCommandType)
	equalTest(t, cmd, "echo flupp\n")
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

func TestBashMultiLine(t *testing.T) {
	c := spawnBash()
	queryTest(t, c, "echo \"multi\nline\"\n", "multi\r\nline\r\n")
}

