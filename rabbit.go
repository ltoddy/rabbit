package rabbit

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

type Handler interface {
	Serve(ctx *Context)
}

type HandlerFunc func(ctx *Context)

func (f HandlerFunc) Serve(ctx *Context) {
	f(ctx)
}

type Rabbit struct {
	Addr   string
	router *router
	*Nest
	nests []*Nest
}

func (r Rabbit) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	p := request.URL.Path
	log.Printf("incmoing request: %-7s %s\n", method, p)

	node, params := r.router.search(method, Path(p))
	if node != nil {
		ctx := newContext(writer, request, params)
		node.handler.Serve(ctx)
	} else {
		_, _ = fmt.Fprintln(writer, "404. That's an error.")
	}
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.router = newRouter()
	rabbit.nests = make([]*Nest, 0, 128)
	rabbit.Nest = NewNest("/", rabbit)
	return rabbit
}

func (r *Rabbit) GenerateNest(prefix Path) *Nest {
	nest := NewNest(prefix, r)
	return nest
}

func (r *Rabbit) register(method string, p Path, handler Handler) {
	if !strings.HasPrefix(string(p), "/") {
		p = Path(path.Join("/", string(p)))
	}

	r.router.register(method, p, handler)
}

func (r *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", r.Addr)
	log.Fatal(http.ListenAndServe(r.Addr, r))
}
