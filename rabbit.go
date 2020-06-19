package rabbit

import (
	"log"
	"net/http"
)

type Rabbit struct {
	Addr   string
	routes map[string]http.HandlerFunc
}

func (r Rabbit) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.routes = make(map[string]http.HandlerFunc)
	return rabbit
}

func (r *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", r.Addr)
	log.Fatal(http.ListenAndServe(r.Addr, r))
}
