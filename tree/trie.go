package tree

import (
	"github.com/ltoddy/rabbit/handler"
	"strings"
)

type TrieTree struct {
	root *trieNode
}

func NewTrieTree() *TrieTree {
	dummy := new(trieNode)
	tree := new(TrieTree)
	tree.root = dummy
	return tree
}

func (tree *TrieTree) Insert(path string, fn handler.Handler) {
	subpaths := splitPath(path)
	tree.root.Insert(path, subpaths, fn, 0)
}

func (tree *TrieTree) Search(path string) (handler.Handler, map[string]string) {
	subpaths := splitPath(path)
	params := make(map[string]string)
	node := tree.root.Search(subpaths, 0)

	if node != nil {
		for index, subpath := range splitPath(node.fullpath) {
			if isDynamicSubPath(subpath) {
				params[interceptDynamicParam(subpath)] = subpaths[index]
			}
		}

		return node.fn, params
	}

	return nil, nil
}

func splitPath(path string) []string {
	subpaths := make([]string, 0)
	for _, subpath := range strings.Split(path, "/") {
		if subpath != "" {
			subpaths = append(subpaths, subpath)
		}
	}
	return subpaths
}
