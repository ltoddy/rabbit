package rabbit

import (
	"net/http"
)

type Nest struct {
	prefix Path
	rabbit *Rabbit
}

func NewNest(prefix Path, rabbit *Rabbit) *Nest {
	nest := new(Nest)
	nest.prefix = prefix
	nest.rabbit = rabbit
	return nest
}

func (n *Nest) Get(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodGet, p, handler)
}

func (n *Nest) Post(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodPost, p, handler)
}

func (n *Nest) Delete(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodDelete, p, handler)
}

func (n *Nest) Patch(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodPatch, p, handler)
}

func (n *Nest) Put(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodPut, p, handler)
}

func (n *Nest) Connect(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodConnect, p, handler)
}

func (n *Nest) Head(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodHead, p, handler)
}

func (n *Nest) Trace(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodTrace, p, handler)
}

func (n *Nest) Options(p Path, handler HandlerFunc) {
	p = n.prefix.Append(p)
	n.rabbit.register(http.MethodOptions, p, handler)
}
