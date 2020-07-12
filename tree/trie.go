package tree

import (
	"github.com/ltoddy/rabbit/handler"
	"strings"
)

// Trie tree
type TrieTree struct {
	root *trieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{root: newTrieNode()}
}

func (tree *TrieTree) Insert(p string, handler handler.Handler) {
	crawler := tree.root
	subpaths := strings.Split(p, "/")
	for _, subpath := range subpaths {
		if _, found := crawler.children[subpath]; !found {
			crawler.children[subpath] = newTrieNode()
		}
		crawler = crawler.children[subpath]
	}
	crawler.end = true
	crawler.handler = handler
}

func (tree *TrieTree) Search(p string) handler.Handler {
	crawler := tree.root
	subpaths := strings.Split(p, "/")
	for _, subpath := range subpaths {
		if _, found := crawler.children[subpath]; !found {
			return nil
		}

		crawler = crawler.children[subpath]
	}

	if crawler.end {
		return crawler.handler
	}

	return nil
}
