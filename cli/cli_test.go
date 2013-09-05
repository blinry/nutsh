package cli

import (
	"testing"
	"time"
)

func spawnBash() CLI {
	return Spawn("bash")
}

func spawnRuby() CLI {
	return Spawn("ruby")
}

func spawnPython() CLI {
	return Spawn("python")
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

	queryTest(t, b, "echo hi", "hi\r\n")
	queryTest(t, b, "echo $((1+1))", "2\r\n")
}

func TestBashStressQueries(t *testing.T) {
	b := spawnBash()

	for i := 0; i<100; i++ {
		queryTest(t, b, "echo hi", "hi\r\n")
	}
}

func TestBashTabCompletion(t *testing.T) {
	b := spawnBash()

	queryTest(t, b, "echo /t\t", "/tmp/\r\n")
	queryTest(t, b, "echo /usr/\t\ti\t", "/usr/include/\r\n")
}

func TestBashEditing(t *testing.T) {
	b := spawnBash()
	queryTest(t, b, "echo notthisfuhi", "hi\r\n")
	queryTest(t, b, "echo worldODODODODODhello ", "hello world\r\n")
}

func TestBashHistory(t *testing.T) {
	b := spawnBash()
	queryTest(t, b, "echo rememberme", "rememberme\r\n")
	queryTest(t, b, "OA", "rememberme\r\n")
}

/*
func TestBashSkipQuery(t *testing.T) {
	b := spawnBash()
	b.send("echo hi\n")
	b.Query("echo something else\n")
	b.send("OA\n")

	equalTest(t, 
	queryTest(t, b, "echo rememberme", "rememberme\r\n")
	queryTest(t, b, "OA", "rememberme\r\n")
}
*/

func TestBashLoop(t *testing.T) {
	c := spawnBash()
	c.allowInteractivity = false

	c.read(promptType)
	c.send("echo flupp\n")
	cmd, _ := c.read(finalCommandType)
	equalTest(t, cmd, "echo flupp\n")
	output, i := c.ReadOutput()
	equalTest(t, output, "flupp\r\n")
	equalTest(t, i, false)
}

func TestBashEmpty(t *testing.T) {
	c := spawnBash()
	queryTest(t, c, "", "")
}

func TestBashMultiLine(t *testing.T) {
	c := spawnBash()
	queryTest(t, c, "echo \"multi\nline\"", "multi\r\nline\r\n")
}

func TestBashInteractive(t *testing.T) {
	c := spawnBash()
	c.read(promptType)
	c.send("vim\n")
	c.read(finalCommandType)
	for {
		select {
		case <-c.runes:
			return
		case <-c.tokens:
			// avoid blocking
		case <-time.After(1*time.Second):
			t.Fatal("No interactivity after 1 second")
		}
	}
}

func TestBashInteractiveAPI(t *testing.T) {
	c := spawnBash()
	c.read(promptType)
	c.send("vim\n")
	go func() {
		<-time.After(500*time.Millisecond)
		c.send(":q\n")
	}()
	_, i := c.ReadOutput()
	equalTest(t, i, true)
}


func TestBashInteractiveAPI(t *testing.T) {
	c := spawnBash()
	c.read(promptType)
	c.send("vim\n")
	go func() {
		<-time.After(500*time.Millisecond)
		c.send(":q\n")
	}()
	_, i := c.ReadOutput()
	equalTest(t, i, true)
}


func TestRubyQueries(t *testing.T) {
	b := spawnRuby()

	queryTest(t, b, "1+1", "2\r\n")
}

func TestPythonQueries(t *testing.T) {
	b := spawnPython()

	queryTest(t, b, "1+1", "2\r\n")
}

func TestPythonMultiLine(t *testing.T) {
	c := spawnPython()
	queryTest(t, c, "if 1:\n\t42\n", "42\r\n")
}
