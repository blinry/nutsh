package cli

import "testing"

func spawnBash() CLI {
    t := target{"bash --norc -i"}
    return Spawn(t)
}

func TestQuery(t *testing.T) {
    b := spawnBash()
    r := b.Query("echo hi")
    if r != "hi\r" {
        t.Fatalf("Expected 'hi\\r'")
    }
    r = b.Query("echo $((1+1))")
    if r != "2\r" {
        t.Fatalf("Expected '2\\r'")
    }
}
