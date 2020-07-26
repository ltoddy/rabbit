package rabbit

import (
	"github.com/ltoddy/rabbit/internal"
	"github.com/ltoddy/rabbit/internal/tree"
	"github.com/ltoddy/rabbit/request"
	"net/http"
)

type router struct {
	// Key: http method
	roots map[string]*tree.TrieTree
}

func newRouter() *router {
	roots := make(map[string]*tree.TrieTree)
	roots[http.MethodGet] = tree.NewTrieTree()
	roots[http.MethodHead] = tree.NewTrieTree()
	roots[http.MethodPost] = tree.NewTrieTree()
	roots[http.MethodPut] = tree.NewTrieTree()
	roots[http.MethodPatch] = tree.NewTrieTree()
	roots[http.MethodDelete] = tree.NewTrieTree()
	roots[http.MethodConnect] = tree.NewTrieTree()
	roots[http.MethodOptions] = tree.NewTrieTree()
	roots[http.MethodTrace] = tree.NewTrieTree()
	return &router{roots}
}

func (router *router) register(method string, path string, handler internal.Handler) {
	root := router.roots[method]

	root.Insert(path, handler)
}

func (router *router) inquiry(method string, path string) (internal.Handler, request.Params) {
	root := router.roots[method]

	return root.Search(path)
}
