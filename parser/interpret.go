package parser

import (
	"morr.cc/nutsh.git/dsl"
	"regexp"
)

type scope struct {
	defs map[string]Node
}

func Interpret(n Node) {
	dsl.Spawn("bash")
	interpret(n, scope{defs: make(map[string]Node)})
}

func interpret(n Node, s scope) string {
	switch n.typ {
	case "block":
		var v string
		for _, node := range(n.children) {
			v = interpret(node, s)
			if v == "break" {
				return "break"
			}
		}
		return v
	case "prompt":
		block := n.children[0]
		for {
			dsl.Prompt()
			if interpret(block, s) == "break" {
				break
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
	case "def":
		name := n.children[0].typ
		s.defs[name] = n
	case "call":
		method := n.children[0].typ
		arguments := n.children[1]
		evaluated_arguments := make([]string, 0)
		for _, arg := range(arguments.children) {
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
		default:
			def, ok := s.defs[method]
			if ok {
				for i, arg := range(def.children[1].children) {
					name := arg.children[0].typ
					s.defs[name] = node("def", node(name), node("arguments"), node("block", node("return", node("string", node(evaluated_arguments[i])))))

				}
				block := def.children[2]
				return interpret(block, s)
			} else {
				panic("Cannot find method '"+method+"'.")
			}
		}
	case "return":
		return interpret(n.children[0], s)
	case "+":
		return interpret(n.children[0], s)+interpret(n.children[1], s)
	case "string":
		return n.children[0].typ
	default:
		panic("I don't know how to interpret a '"+n.typ+"' node.")
	}
	return "whatever"
}
