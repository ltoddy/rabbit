package tree

import (
	"github.com/ltoddy/rabbit/handler"
)

type trieNode struct {
	// sorted, and use binary search
	children map[string]*trieNode

	end     bool // end is true if node represents end of a path
	handler handler.Handler
}

func newTrieNode() *trieNode {
	node := new(trieNode)
	node.children = make(map[string]*trieNode)
	node.end = false
	node.handler = nil
	return node
}
