package tree

import (
	"github.com/ltoddy/rabbit"
	"net/http"
	"testing"
)

type TestHandler struct{}

func (t TestHandler) Serve(request *http.Request) rabbit.Response {
	return rabbit.TextResponse("hello world")
}

func TestTrieTree(t *testing.T) {
	t.Run("should get handler", func(t *testing.T) {
		tree := NewTrieTree()
		handler := TestHandler{}

		tree.Insert("/hello/world", handler)

		if handler != tree.Search("/hello/world") {
			t.Fail()
		}
	})

	t.Run("should get nil given an empty trie tree", func(t *testing.T) {
		tree := NewTrieTree()

		if tree.Search("/hello/world") != nil {
			t.Fail()
		}
	})
}
