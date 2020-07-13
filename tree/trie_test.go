package tree

import (
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
	"testing"
)

type TestHandler struct{}

func (t TestHandler) Serve(request *request.Request) response.Response {
	return response.TextResponse("hello world")
}

func TestTrieTree(t *testing.T) {
	t.Run("should get handler", func(t *testing.T) {
		tree := NewTrieTree()
		handler := TestHandler{}

		tree.Insert("/hello/world", handler)

		actual, _ := tree.Search("/hello/world")

		if handler != actual {
			t.Fail()
		}
	})

	t.Run("should get nil given an empty trie tree", func(t *testing.T) {
		tree := NewTrieTree()

		actual, _ := tree.Search("/hello/world")
		if actual != nil {
			t.Fail()
		}
	})
}

func TestIsDynamicSubPath(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		subpath := "<name>"

		if isDynamicSubPath(subpath) == false {
			t.Fail()
		}
	})

	t.Run("should return false", func(t *testing.T) {
		subpaths := []string{"<<name>", "some", "<123sada"}

		for _, subpath := range subpaths {
			if isDynamicSubPath(subpath) {
				t.Fail()
			}
		}
	})
}

func TestInterceptDynamicParam(t *testing.T) {
	t.Run("should return both side angle brackets ", func(t *testing.T) {
		subpath := "<name>"
		expected := "name"

		actual := interceptDynamicParam(subpath)

		if expected != actual {
			t.Errorf("expected: %s, got: %s\n", expected, actual)
		}

	})
}
