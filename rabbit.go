package rabbit

import (
	"fmt"
	"github.com/ltoddy/rabbit/types"
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
	routes map[string]Handler
	*Nest
	nests []*Nest
}

func (r Rabbit) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	p := request.URL.Path
	key := fmt.Sprintf("%s-%s", method, p)
	log.Printf("incmoing request: %-7s %s\n", method, p)

	handler, ok := r.routes[key]
	ctx := NewContext(writer, request)
	if ok {
		handler.Serve(ctx)
	} else {
		_, _ = fmt.Fprintln(writer, "404. That's an error.")
	}
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.routes = make(map[string]Handler)
	rabbit.nests = make([]*Nest, 0, 128)
	rabbit.Nest = NewNest("/", rabbit)
	return rabbit
}

func (r *Rabbit) GenerateNest(prefix types.Path) *Nest {
	nest := NewNest(prefix, r)
	return nest
}

func (r *Rabbit) register(method string, p types.Path, handler Handler) {
	if !strings.HasPrefix(string(p), "/") {
		p = types.Path(path.Join("/", string(p)))
	}

	key := fmt.Sprintf("%s-%s", method, p)
	r.routes[key] = handler
}

func (r *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", r.Addr)
	log.Fatal(http.ListenAndServe(r.Addr, r))
}
