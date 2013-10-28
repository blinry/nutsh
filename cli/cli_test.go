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
	o, ok := c.Query(command)
	equalTest(t, o, output)
	equalTest(t, ok, true)
}

func TestBashQueries(t *testing.T) {
	b := spawnBash()
	defer b.Quit()

	queryTest(t, b, "echo hi", "hi\r\n")
	queryTest(t, b, "echo $((1+1))", "2\r\n")
}

func TestBashStressQueries(t *testing.T) {
	b := spawnBash()
	defer b.Quit()

	for i := 0; i < 100; i++ {
		queryTest(t, b, "echo hi", "hi\r\n")
	}
}

func TestBashTabCompletion(t *testing.T) {
	b := spawnBash()
	defer b.Quit()

	queryTest(t, b, "echo /t\t", "/tmp/\r\n")
	queryTest(t, b, "echo /usr/\t\ti\t", "/usr/include/\r\n")
}

func TestBashEditing(t *testing.T) {
	b := spawnBash()
	defer b.Quit()
	queryTest(t, b, "echo notthisfuhi", "hi\r\n")
	queryTest(t, b, "echo worldODODODODODhello ", "hello world\r\n")
}

func TestBashHistory(t *testing.T) {
	b := spawnBash()
	defer b.Quit()
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
	defer c.Quit()
	c.allowInteractivity = false

	c.read(promptType)
	c.send("echo flupp\n")
	cmd, _, _ := c.read(finalCommandType)
	equalTest(t, cmd, "echo flupp\n")
	output, i, _ := c.ReadOutput()
	equalTest(t, output, "flupp\r\n")
	equalTest(t, i, false)
}

func TestBashEmpty(t *testing.T) {
	c := spawnBash()
	defer c.Quit()
	queryTest(t, c, "", "")
}

func TestBashMultiLine(t *testing.T) {
	c := spawnBash()
	defer c.Quit()
	queryTest(t, c, "echo \"multi\nline\"", "multi\r\nline\r\n")
}

func TestBashInteractive(t *testing.T) {
	c := spawnBash()
	defer c.Quit()
	c.read(promptType)
	c.send("vim\n")
	c.read(finalCommandType)
	for {
		select {
		case <-c.runes:
			return
		case <-c.tokens:
			// avoid blocking
		case <-time.After(1 * time.Second):
			t.Fatal("No interactivity after 1 second")
		}
	}
}

func TestBashInteractiveAPI(t *testing.T) {
	c := spawnBash()
	defer c.Quit()
	c.read(promptType)
	c.send("vim\n")
	go func() {
		<-time.After(500 * time.Millisecond)
		c.send(":q\n")
	}()
	_, i, _ := c.ReadOutput()
	equalTest(t, i, true)
}

func TestBashQueryInteractive(t *testing.T) {
	c := spawnBash()
	defer c.Quit()

	c.QueryInteractive("man man", "q")
	queryTest(t, c, "echo hi", "hi\r\n")
}

func TestRubyQueries(t *testing.T) {
	b := spawnRuby()
	defer b.Quit()

	queryTest(t, b, "1+1", "2\r\n")
}

/*
func TestPythonQueries(t *testing.T) {
	b := spawnPython()
	defer b.Quit()

	queryTest(t, b, "1+1", "2\r\n")
}

func TestPythonMultiLine(t *testing.T) {
	c := spawnPython()
	defer c.Quit()
	queryTest(t, c, "if 1:\n\t42\n", "42\r\n")
}
*/
