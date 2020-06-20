package rabbit

import (
	"github.com/ltoddy/rabbit/types"
	"net/http"
)

type Nest struct {
	prefix types.Path
	rabbit *Rabbit
}

func NewNest(prefix types.Path, rabbit *Rabbit) *Nest {
	nest := new(Nest)
	nest.prefix = prefix
	nest.rabbit = rabbit
	return nest
}

func (n *Nest) Get(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodGet, p, handler)
}

func (n *Nest) Post(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodPost, p, handler)
}

func (n *Nest) Delete(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodDelete, p, handler)
}

func (n *Nest) Patch(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodPatch, p, handler)
}

func (n *Nest) Put(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodPut, p, handler)
}

func (n *Nest) Connect(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodConnect, p, handler)
}

func (n *Nest) Head(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodHead, p, handler)
}

func (n *Nest) Trace(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodTrace, p, handler)
}

func (n *Nest) Options(p types.Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodOptions, p, handler)
}
