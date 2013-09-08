package parser

import (
	"fmt"
)

func Test(n Node) {
	expects := annotate(&n)
	fmt.Println(len(expects))
	fmt.Println(expects)
	fmt.Println(n)
	InterpretTest(n)
}

func annotate(n *Node) []Node {
	expects := make([]Node, 0)
	for i := range(n.children) {
		c := &n.children[i]
		if c.typ == "prompt" {
			e := collect_expects(*c)
			c.children = append(c.children, node("excpects", e...))
			fmt.Println(c)
			expects = append(expects, e...)
		}
		expects = append(expects, annotate(c)...)
	}
	return expects
}

func collect_expects(n Node) []Node {
	expects := make([]Node, 0)
	for _, c := range(n.children) {
		if c.typ == "call" {
			if c.children[0].typ == "expect" {
				expects = append(expects, c.children[1].children...)
			}
		}
		if c.typ != "prompt" {
			expects = append(expects, collect_expects(c)...)
		}
	}
	return expects
}
