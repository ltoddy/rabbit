package rabbit

import (
	"fmt"
	"log"
	"net/http"
)

type Handler interface {
	Serve(*http.Request) Response
}

type HandlerFunc func(*http.Request) Response

func (f HandlerFunc) Serve(r *http.Request) Response {
	return f(r)
}

type Rabbit struct {
	Addr   string
	router map[string]Handler
}

func (r Rabbit) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	p := request.URL.Path
	log.Printf("incmoing request: %-7s %s\n", method, p)
	key := fmt.Sprintf("%s-%s", method, p)

	handler, ok := r.router[key]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	response := handler.Serve(request)

	writer.WriteHeader(response.StatusCode())
	for key, value := range response.Header() {
		writer.Header().Set(key, value)
	}
	_, _ = writer.Write(response.Body())
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.router = make(map[string]Handler)
	return rabbit
}

func (r *Rabbit) Get(p string, f HandlerFunc) {
	key := fmt.Sprintf("%s-%s", http.MethodGet, p)
	r.router[key] = f
}

func (r *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", r.Addr)
	log.Fatal(http.ListenAndServe(r.Addr, r))
}
