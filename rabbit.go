package rabbit

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

type Rabbit struct {
	Addr   string
	routes map[string]http.HandlerFunc
	*Nest
	nests []*Nest
}

func (r Rabbit) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	p := request.URL.Path
	key := fmt.Sprintf("%s-%s", method, p)
	log.Printf("accept request: %s\n", p)

	handler, ok := r.routes[key]
	if ok {
		handler(writer, request)
	} else {
		_, _ = fmt.Fprintln(writer, "404. That's an error.")
	}
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.routes = make(map[string]http.HandlerFunc)
	rabbit.nests = make([]*Nest, 0, 128)
	rabbit.Nest = NewNest("/", rabbit)
	return rabbit
}

func (r *Rabbit) GenerateNest(prefix Path) *Nest {
	nest := NewNest(prefix, r)
	return nest
}

func (r *Rabbit) register(method string, p Path, handler http.HandlerFunc) {
	if !strings.HasPrefix(string(p), "/") {
		p = Path(path.Join("/", string(p)))
	}

	key := fmt.Sprintf("%s-%s", method, p)
	r.routes[key] = handler
}

func (r *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", r.Addr)
	log.Fatal(http.ListenAndServe(r.Addr, r))
}
