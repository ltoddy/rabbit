package tree

import (
	"github.com/ltoddy/rabbit/handler"
	"regexp"
	"strings"
)

type trieNode struct {
	fullpath string
	subpath  string
	children []*trieNode
	dynamic  bool
	fn       handler.Handler
}

func NewTrieNode(fullpath string, subpath string, fn handler.Handler) *trieNode {
	dynamic := isDynamicSubPath(subpath)

	return &trieNode{
		fullpath: fullpath,
		subpath:  subpath,
		children: make([]*trieNode, 0),
		dynamic:  dynamic,
		fn:       fn,
	}
}

func (node *trieNode) Insert(fullpath string, subpaths []string, fn handler.Handler, level int) {
	if len(subpaths) == level {
		node.fullpath = fullpath
		return
	}

	subpath := subpaths[level]
	child := node.matchChild(subpath)
	if child == nil {
		child = NewTrieNode(fullpath, subpath, fn)
		node.children = append(node.children, child)
	}
	child.Insert(fullpath, subpaths, fn, level+1)
}

func (node *trieNode) Search(subpaths []string, level int) *trieNode {
	if len(subpaths) == level || node.dynamic {
		if node.subpath == "" {
			return nil
		}
		return node
	}

	subpath := subpaths[level]
	children := node.matchChildren(subpath)

	for _, child := range children {
		result := child.Search(subpaths, level+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (node *trieNode) matchChild(subpath string) *trieNode {
	for _, child := range node.children {
		if child.subpath == subpath || child.dynamic {
			return child
		}
	}

	return nil
}

func (node *trieNode) matchChildren(subpath string) []*trieNode {
	children := make([]*trieNode, 0)
	for _, child := range node.children {
		if child.subpath == subpath || child.dynamic {
			children = append(children, child)
		}
	}
	return children
}

func isDynamicSubPath(subpath string) bool {
	return regexp.MustCompile(`^<\w+>$`).MatchString(subpath)
}

func interceptDynamicParam(subpath string) string {
	subpath = strings.TrimPrefix(subpath, "<")
	subpath = strings.TrimSuffix(subpath, ">")
	return subpath
}
