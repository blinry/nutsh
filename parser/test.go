package parser

import (
	//"fmt"
)

func Test(n *Node, common *Node) {
	expects := annotate(n)

	repeat:
	unreachedExpects := 0
	InterpretTest(n, common)
	for _, e := range(expects) {
		if e.children[2].children[0].typ == "false" {
			unreachedExpects += 1
		}
	}
	if unreachedExpects > 0 {
		goto repeat
	}
}

func annotate(n *Node) []*Node {
	expects := make([]*Node, 0)
	for i := range n.children {
		c := n.children[i]
		if c.typ == "prompt" {
			e := collect_expects(c)
			c.children = append(c.children, node("excpects", e...))
			//fmt.Println(c)
			expects = append(expects, e...)
		}
		expects = append(expects, annotate(c)...)
	}
	return expects
}

func collect_expects(n *Node) []*Node {
	expects := make([]*Node, 0)
	for _, c := range n.children {
		if c.typ == "call" {
			if c.children[0].typ == "expect" {
				c.children = append(c.children, node("reached", node("false")))
				expects = append(expects, c)
			}
		}
		if c.typ != "prompt" {
			expects = append(expects, collect_expects(c)...)
		}
	}
	return expects
}
