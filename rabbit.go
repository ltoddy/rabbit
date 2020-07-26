package rabbit

import (
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
	"log"
	"net/http"
)

type Rabbit struct {
	Addr       string
	Router     *router
	blueprints map[string]*Blueprint
}

func NewRabbit(addr string) *Rabbit {
	rabbit := new(Rabbit)
	rabbit.Addr = addr
	rabbit.Router = newRouter()
	return rabbit
}

func (rabbit *Rabbit) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	method := r.Method
	p := r.URL.Path
	log.Printf("incmoing request: %-7s %s\n", method, p)

	handle, params := rabbit.Router.inquiry(method, p)
	if handle == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	req := request.NewRequest(r, params)
	defer recovery(writer)
	resp := handle.Serve(req)

	h := writer.Header()
	for key, value := range resp.Header() {
		h.Set(key, value)
	}
	writer.WriteHeader(resp.StatusCode())
	_, _ = writer.Write(resp.Body())
}

func (rabbit *Rabbit) Get(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodGet, path, f)
}

func (rabbit *Rabbit) Head(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodHead, path, f)
}

func (rabbit *Rabbit) Post(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodPost, path, f)
}

func (rabbit *Rabbit) Put(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodPut, path, f)
}

func (rabbit *Rabbit) Patch(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodPatch, path, f)
}

func (rabbit *Rabbit) Delete(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodDelete, path, f)
}

func (rabbit *Rabbit) Connect(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodConnect, path, f)
}

func (rabbit *Rabbit) Options(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodOptions, path, f)
}

func (rabbit *Rabbit) Trace(path string, f HandlerFunction) {
	rabbit.Router.register(http.MethodTrace, path, f)
}

func (rabbit *Rabbit) RegisterBlueprint(blueprint *Blueprint) {
	for _, entry := range blueprint.get {
		rabbit.Router.register(http.MethodGet, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.head {
		rabbit.Router.register(http.MethodHead, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.post {
		rabbit.Router.register(http.MethodPost, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.put {
		rabbit.Router.register(http.MethodPut, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.patch {
		rabbit.Router.register(http.MethodPatch, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.delete {
		rabbit.Router.register(http.MethodDelete, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.connect {
		rabbit.Router.register(http.MethodConnect, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.options {
		rabbit.Router.register(http.MethodOptions, entry.fullpath, entry.fn)
	}
	for _, entry := range blueprint.trace {
		rabbit.Router.register(http.MethodTrace, entry.fullpath, entry.fn)
	}
}

func (rabbit *Rabbit) RegisterBlueprints(blueprints ...*Blueprint) {
	for _, blueprint := range blueprints {
		rabbit.RegisterBlueprint(blueprint)
	}
}

func (rabbit *Rabbit) Run() {
	log.Printf("Server start run at: %s\n", rabbit.Addr)
	log.Fatal(http.ListenAndServe(rabbit.Addr, rabbit))
}

type HandlerFunction func(*request.Request) response.Response

func (f HandlerFunction) Serve(r *request.Request) response.Response {
	return f(r)
}
