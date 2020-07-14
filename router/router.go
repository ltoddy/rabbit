package router

import (
	"github.com/ltoddy/rabbit/handler"
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/tree"
	"net/http"
)

type Router struct {
	// Key: http method
	roots map[string]*tree.TrieTree
}

func NewRouter() *Router {
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
	return &Router{roots}
}

func (router *Router) Register(method string, path string, handler handler.Handler) {
	root := router.roots[method]

	root.Insert(path, handler)
}

func (router *Router) Inquiry(method string, path string) (handler.Handler, request.Params) {
	root := router.roots[method]

	return root.Search(path)
}
