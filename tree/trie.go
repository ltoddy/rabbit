package tree

import (
	"github.com/ltoddy/rabbit/handler"
	"github.com/ltoddy/rabbit/request"
	"regexp"
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

func (tree *TrieTree) Search(p string) (handler.Handler, request.Params) {
	crawler := tree.root
	params := make(request.Params)
	subpaths := strings.Split(p, "/")
	for _, subpath := range subpaths {
		if _, found := crawler.children[subpath]; !found {
			return nil, nil
		}

		crawler = crawler.children[subpath]
	}

	if crawler.end {
		return crawler.handler, params
	}

	return nil, nil
}

func isDynamicSubPath(subpath string) bool {
	return regexp.MustCompile(`^<\w+>$`).MatchString(subpath)
}

func interceptDynamicParam(subpath string) string {
	subpath = strings.TrimPrefix(subpath, "<")
	subpath = strings.TrimSuffix(subpath, ">")
	return subpath
}
