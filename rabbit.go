package rabbit

import (
	"github.com/ltoddy/rabbit/Router"
	"github.com/ltoddy/rabbit/handler"
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

func (r *Rabbit) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	p := request.URL.Path
	log.Printf("incmoing request: %-7s %s\n", method, p)

	handle := r.Router.Inquiry(method, p)
	if handle == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	response := handle.Serve(request)

	writer.WriteHeader(response.StatusCode())
	for key, value := range response.Header() {
		writer.Header().Set(key, value)
	}
	_, _ = writer.Write(response.Body())
}

func (r *Rabbit) Get(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodGet, path, f)
}

func (r *Rabbit) Head(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodHead, path, f)
}

func (r *Rabbit) Post(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodPost, path, f)
}

func (r *Rabbit) Put(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodPut, path, f)
}

func (r *Rabbit) Patch(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodPatch, path, f)
}

func (r *Rabbit) Delete(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodDelete, path, f)
}

func (r *Rabbit) Connect(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodConnect, path, f)
}

func (r *Rabbit) Options(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodOptions, path, f)
}

func (r *Rabbit) Trace(path string, f handler.HandlerFunction) {
	r.Router.Register(http.MethodTrace, path, f)
}

func (r *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", r.Addr)
	log.Fatal(http.ListenAndServe(r.Addr, r))
}
