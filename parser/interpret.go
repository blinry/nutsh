package parser

import (
	"morr.cc/nutsh.git/dsl"
	"regexp"
)

type scope struct {
	defs           map[string]Node
	blocks         []Node
	test           bool
	current_expect string
}

func Interpret(n Node) {
	dsl.Spawn("bash")
	interpret(n, &scope{defs: make(map[string]Node), blocks: make([]Node, 0), test: false})
}

func InterpretTest(n Node) {
	dsl.Spawn("bash")
	interpret(n, &scope{defs: make(map[string]Node), blocks: make([]Node, 0), test: true})
}

func interpret(n Node, s *scope) string {
	switch n.typ {
	case "block":
		var v string
		for _, node := range n.children {
			v = interpret(node, s)
			if v == "break" {
				return "break"
			}
		}
		return v
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
				dsl.Prompt()
			}
			for _, block := range s.blocks {
				interpret(block, s)
			}
			if interpret(block, s) == "break" {
				break
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
		if interpret(condition, s) == "" {
			if interpret(else_block, s) == "break" {
				return "break"
			}
		} else {
			if interpret(block, s) == "break" {
				return "break"
			}
		}
	case "state":
		promptblock := n.children[0]
		s.blocks = append(s.blocks, promptblock)
		block := n.children[1]
		interpret(block, s)
	case "def":
		name := n.children[0].typ
		s.defs[name] = n
	case "call":
		method := n.children[0].typ
		arguments := n.children[1]
		evaluated_arguments := make([]string, 0)
		for _, arg := range arguments.children {
			evaluated_arguments = append(evaluated_arguments, interpret(arg, s))
		}

		switch method {
		case "say":
			dsl.Say(evaluated_arguments[0])
		case "command":
			return dsl.LastCommand()
		case "output":
			return dsl.LastOutput()
		case "match":
			if regexp.MustCompile(evaluated_arguments[1]).MatchString(evaluated_arguments[0]) {
				return "true"
			} else {
				return ""
			}
		case "equal":
			if evaluated_arguments[0] == evaluated_arguments[1] {
				return "true"
			} else {
				return ""
			}
		case "run":
			return dsl.Query(evaluated_arguments[0])
		case "break":
			return "break"
		case "return":
			return evaluated_arguments[0]
		case "expect":
			if evaluated_arguments[0] == s.current_expect {
				s.current_expect = ""
			}
			return ""
		default:
			def, ok := s.defs[method]
			if ok {
				for i, arg := range def.children[1].children {
					name := arg.children[0].typ
					s.defs[name] = node("def", node(name), node("arguments"), node("block", node("call", node("return"), node("stringexpressions", node("string", node(evaluated_arguments[i]))))))

				}
				block := def.children[2]
				return interpret(block, s)
			} else {
				panic("Cannot find method '" + method + "'.")
			}
		}
	case "+":
		return interpret(n.children[0], s) + interpret(n.children[1], s)
	case "string":
		return n.children[0].typ
	case "and":
		return bool2str(str2bool(interpret(n.children[0], s)) && str2bool(interpret(n.children[1], s)))
	case "or":
		return bool2str(str2bool(interpret(n.children[0], s)) || str2bool(interpret(n.children[1], s)))
	case "not":
		return bool2str(!str2bool(interpret(n.children[0], s)))
	default:
		panic("I don't know how to interpret a '" + n.typ + "' node.")
	}
	return "whatever"
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
