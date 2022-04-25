package gee

import (
	"fmt"
)

type node struct {
	pattern   string
	part      string
	children  []*node
	isWild    bool
	isGeneral bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%v, part=%v, isWild=%v, children =%v}", n.pattern, n.part, n.isWild, n.children)
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild || child.isGeneral {
			return child
		}
	}
	return nil
}

// parts 传递下去，减少计算
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height || n.isGeneral {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':', isGeneral: part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild || child.isGeneral {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || n.isGeneral {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
