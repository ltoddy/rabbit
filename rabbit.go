package rabbit

import (
	"fmt"
	"log"
	"net/http"
)

type Rabbit struct {
	Addr   string
	routes map[string]http.HandlerFunc
}

func (r Rabbit) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	path := request.URL.Path
	key := fmt.Sprintf("%s-%s", method, path)
	log.Printf("accept request: %s\n", path)

	handler, ok := r.routes[key]
	if ok {
		handler(writer, request)
	} else {
		_, _ = fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL.Path)
	}
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.routes = make(map[string]http.HandlerFunc)
	return rabbit
}

func (r *Rabbit) register(method string, path Path, handler http.HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, path)
	r.routes[key] = handler
}

func (r *Rabbit) Get(path Path, handler http.HandlerFunc) {
	r.register(http.MethodGet, path, handler)
}

func (r *Rabbit) Post(path Path, handler http.HandlerFunc) {
	r.register(http.MethodPost, path, handler)
}

func (r *Rabbit) Delete(path Path, handler http.HandlerFunc) {
	r.register(http.MethodDelete, path, handler)
}

func (r *Rabbit) Patch(path Path, handler http.HandlerFunc) {
	r.register(http.MethodPatch, path, handler)
}

func (r *Rabbit) Put(path Path, handler http.HandlerFunc) {
	r.register(http.MethodPut, path, handler)
}

func (r *Rabbit) Connect(path Path, handler http.HandlerFunc) {
	r.register(http.MethodConnect, path, handler)
}

func (r *Rabbit) Head(path Path, handler http.HandlerFunc) {
	r.register(http.MethodHead, path, handler)
}

func (r *Rabbit) Trace(path Path, handler http.HandlerFunc) {
	r.register(http.MethodTrace, path, handler)
}

func (r *Rabbit) Options(path Path, handler http.HandlerFunc) {
	r.register(http.MethodOptions, path, handler)
}

func (r *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", r.Addr)
	log.Fatal(http.ListenAndServe(r.Addr, r))
}
