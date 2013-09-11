package parser

import (
	"morr.cc/nutsh.git/dsl"
	"regexp"
	"time"
)

type scope struct {
	defs           map[string]Node
	blocks         []Node
	test           bool
	current_expect string
}

type interrupt struct {
	typ string
	value string
}

func Interpret(n Node) (string, bool) {
	dsl.Spawn("bash")
	_, i := interpret(n, &scope{defs: make(map[string]Node), blocks: make([]Node, 0), test: false})
	dsl.Quit()
	time.Sleep(1000*time.Millisecond)

	if i.typ == "lesson" {
		return i.value, false
	} else if i.typ == "done" {
		return i.value, true
	} else {
		return "", false
	}
}

func InterpretTest(n Node) {
	dsl.Spawn("bash")
	interpret(n, &scope{defs: make(map[string]Node), blocks: make([]Node, 0), test: true})
}

func interpret(n Node, s *scope) (string, interrupt) {
	i := interrupt{}
	switch n.typ {
	case "lesson":
		_, i = interpret(n.children[0], s)
		if i.typ != "" {
			return "", i
		}
		return "", interrupt{"done", ""}
	case "block":
		var v string
		for _, node := range n.children {
			v, i = interpret(node, s)
			if i.typ != "" {
				return "", i
			}
		}
		return v, i
	case "prompt":
		block := n.children[0]
		expects := node("expects")
		if len(n.children) > 1 {
			expects = n.children[1]
		}
		for {
			if s.test {
				if len(expects.children) > 0 {
					expect := expects.children[0].children[0].typ
					s.current_expect = expect
					dsl.SimulatePrompt(expect)
				} else {
					panic("No expect in prompt")
				}
			} else {
				if ! dsl.Prompt() {
					// cli terminated
					return "", interrupt{"lesson", ""}
				}
			}
			for _, block := range s.blocks {
				_, i := interpret(block, s)
				if i.typ != "" {
					return "", i
				}
			}
			_, i := interpret(block, s)
			if i.typ == "break" {
				break
			}
			if i.typ != "" {
				return "", i
			}
			if s.test {
				if s.current_expect != "" {
					panic("Expect was not reached: " + s.current_expect)
				}
			}
		}
	case "if":
		condition := n.children[0]
		block := n.children[1]
		else_block := n.children[2]
		v, i := interpret(condition, s)
		if i.typ != "" {
			return "", i
		}
		if v == "" {
			_, i = interpret(else_block, s)
		} else {
			_, i = interpret(block, s)
		}
		if i.typ != "" {
			return "", i
		}
	case "state":
		promptblock := n.children[0]
		s.blocks = append(s.blocks, promptblock)
		block := n.children[1]
		_, i := interpret(block, s)
		if i.typ != "" {
			return "", i
		}
	case "def":
		name := n.children[0].typ
		s.defs[name] = n
	case "call":
		method := n.children[0].typ
		arguments := n.children[1]
		evaluated_arguments := make([]string, 0)
		for _, arg := range arguments.children {
			v, i := interpret(arg, s)
			if i.typ != "" {
				return "", i
			}
			evaluated_arguments = append(evaluated_arguments, v)
		}

		switch method {
		case "say":
			dsl.Say(evaluated_arguments[0])
		case "command":
			return dsl.LastCommand(), i
		case "output":
			return dsl.LastOutput(), i
		case "match":
			if regexp.MustCompile(evaluated_arguments[1]).MatchString(evaluated_arguments[0]) {
				return "true", i
			} else {
				return "", i
			}
		case "equal":
			if evaluated_arguments[0] == evaluated_arguments[1] {
				return "true", i
			} else {
				return "", i
			}
		case "run":
			return dsl.Query(evaluated_arguments[0]), i
		case "break":
			return "", interrupt{"break", ""}
		case "lesson":
			return "", interrupt{"lesson", evaluated_arguments[0]}
		case "done":
			return "", interrupt{"done", ""}
		case "return":
			return evaluated_arguments[0], i
		case "expect":
			if evaluated_arguments[0] == s.current_expect {
				s.current_expect = ""
			}
			return "", i
		default:
			def, ok := s.defs[method]
			if ok {
				for i, arg := range def.children[1].children {
					name := arg.children[0].typ
					s.defs[name] = node("def", node(name), node("arguments"), node("block", node("call", node("return"), node("stringexpressions", node("string", node(evaluated_arguments[i]))))))

				}
				block := def.children[2]
				v, i := interpret(block, s)
				if i.typ != "" {
					return "", i
				}
				return v, i
			} else {
				panic("Cannot find method '" + method + "'.")
			}
		}
	case "+":
		v1, i := interpret(n.children[0], s)
		if i.typ != "" {
			return "", i
		}
		v2, i := interpret(n.children[1], s)
		if i.typ != "" {
			return "", i
		}
		return v1+v2, i
	case "string":
		return n.children[0].typ, i
	case "and":
		v1, i := interpret(n.children[0], s)
		if i.typ != "" {
			return "", i
		}
		v2, i := interpret(n.children[1], s)
		if i.typ != "" {
			return "", i
		}
		return bool2str(str2bool(v1) && str2bool(v2)), i
	case "or":
		v1, i := interpret(n.children[0], s)
		if i.typ != "" {
			return "", i
		}
		v2, i := interpret(n.children[1], s)
		if i.typ != "" {
			return "", i
		}
		return bool2str(str2bool(v1) || str2bool(v2)), i
	case "not":
		v, i := interpret(n.children[0], s)
		if i.typ != "" {
			return "", i
		}
		return bool2str(!str2bool(v)), i
	default:
		panic("I don't know how to interpret a '" + n.typ + "' node.")
	}
	return "whatever", i
}

func str2bool(s string) bool {
	return s != ""
}

func bool2str(b bool) string {
	if b {
		return "true"
	} else {
		return ""
	}
}
