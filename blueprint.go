package rabbit

import (
	"github.com/ltoddy/rabbit/handler"
	"github.com/ltoddy/rabbit/router"
	"net/http"
	"path"
)

type Blueprint struct {
	prefix string
	router *router.Router
}

// TODO: how to merge two trie trees
func NewBlueprint(prefix string) *Blueprint {
	if prefix == "" || prefix == "/" {
		panic(`blueprint's prefix should not be empty or "/"`)
	}

	return &Blueprint{
		prefix: prefix,
		router: router.NewRouter(),
	}
}

func (nest *Blueprint) Get(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodGet, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Head(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodHead, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Post(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodPost, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Put(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodPut, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Patch(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodPatch, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Delete(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodDelete, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Connect(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodConnect, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Options(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodOptions, path.Join(nest.prefix, p), fn)
}

func (nest *Blueprint) Trace(p string, fn handler.HandlerFunction) {
	nest.router.Register(http.MethodTrace, path.Join(nest.prefix, p), fn)
}
