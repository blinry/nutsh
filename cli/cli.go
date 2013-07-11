package cli

type CLI struct {
}

func Spawn(t target) CLI {
	return CLI{}
}

func (c CLI) Query(cmd string) string {
	return "hi\r"
}
