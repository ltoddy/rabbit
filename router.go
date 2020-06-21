package rabbit

import (
	"strings"
)

type router struct {
	// Key: http method
	roots map[string]*trieNode
}

func newRouter() *router {
	return &router{roots: make(map[string]*trieNode)}
}

func (router *router) register(method string, p Path, handler Handler) {
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &trieNode{
			children: make([]*trieNode, 0),
		}
	}
	router.roots[method].insert(p, 0, handler)
}

func (router *router) search(method string, p Path) (*trieNode, map[string]string) {
	parts := toparts(p)
	params := make(map[string]string)

	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}

	node := root.search(parts, 0)
	if node != nil {
		searchparts := toparts(node.p)

		for i, part := range searchparts {
			if strings.HasPrefix(part, ":") {
				params[part[1:]] = parts[i]
			}
		}

		return node, params
	}

	return nil, nil
}

type trieNode struct {
	p        Path   // 当前节点的路径 e.g.  /hello/:name/foo
	part     string // 每层节点
	children []*trieNode
	handler  Handler // 叶子节点存放handler,在Search的时候，返回叶子节点
}

func (node *trieNode) matchChild(part string) *trieNode {
	for _, child := range node.children {
		if child.part == part {
			return child
		}
	}

	return nil
}

func (node *trieNode) matchChildren(part string) []*trieNode {
	nodes := make([]*trieNode, 0, len(node.children))
	for _, child := range node.children {
		if child.part == part {
			nodes = append(nodes, child)
		}
		if strings.HasPrefix(child.part, ":") {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (node *trieNode) insert(p Path, level int, handler Handler) {
	parts := toparts(p)
	if len(parts) == level {
		node.handler = handler
		node.p = p
		return
	}

	part := parts[level]
	child := node.matchChild(part)
	if child == nil {
		child = &trieNode{part: part}
		node.children = append(node.children, child)
	}
	child.insert(p, level+1, handler)
}

func (node *trieNode) search(parts []string, level int) *trieNode {
	if len(parts) == level {
		return node
	}

	part := parts[level]
	children := node.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, level+1)

		if result != nil {
			return result
		}
	}

	return nil
}

func toparts(p Path) []string {
	xs := strings.Split(string(p), "/")

	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		if x != "" {
			parts = append(parts, x)
		}
	}

	return parts
}
