package rabbit

import (
	"github.com/ltoddy/rabbit/handler"
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/router"
	"log"
	"net/http"
)

type Rabbit struct {
	Addr   string
	Router *router.Router
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.Router = router.NewRouter()
	return rabbit
}

func (rabbit *Rabbit) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	method := r.Method
	p := r.URL.Path
	log.Printf("incmoing request: %-7s %s\n", method, p)

	handle, _ := rabbit.Router.Inquiry(method, p)
	if handle == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	req := &request.Request{
		Request: r,
		Params:  nil,
	}

	response := handle.Serve(req)

	writer.WriteHeader(response.StatusCode())
	for key, value := range response.Header() {
		writer.Header().Set(key, value)
	}
	_, _ = writer.Write(response.Body())
}

func (rabbit *Rabbit) Get(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodGet, path, f)
}

func (rabbit *Rabbit) Head(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodHead, path, f)
}

func (rabbit *Rabbit) Post(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodPost, path, f)
}

func (rabbit *Rabbit) Put(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodPut, path, f)
}

func (rabbit *Rabbit) Patch(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodPatch, path, f)
}

func (rabbit *Rabbit) Delete(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodDelete, path, f)
}

func (rabbit *Rabbit) Connect(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodConnect, path, f)
}

func (rabbit *Rabbit) Options(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodOptions, path, f)
}

func (rabbit *Rabbit) Trace(path string, f handler.HandlerFunction) {
	rabbit.Router.Register(http.MethodTrace, path, f)
}

func (rabbit *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", rabbit.Addr)
	log.Fatal(http.ListenAndServe(rabbit.Addr, rabbit))
}
