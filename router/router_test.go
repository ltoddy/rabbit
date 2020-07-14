package router

import (
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
	"net/http"
	"testing"
)

type handle struct{}

func (h handle) Serve(r *request.Request) response.Response {
	return response.TextResponse("")
}

func TestStaticRouter(t *testing.T) {
	h := handle{}

	t.Run("should get handler when inquiry for router given correct http method and path", func(t *testing.T) {
		router := NewRouter()
		router.Register(http.MethodGet, "/hello/world", h)

		actual, params := router.Inquiry(http.MethodGet, "/hello/world")

		if actual != h || len(params) != 0 {
			t.Errorf("handler: expected handler, got %#v\nparams: expected {}, got %#v\n", actual, params)
		}
	})

	t.Run("should get nothing when inquiry for router given incorrect http method", func(t *testing.T) {
		router := NewRouter()
		router.Register(http.MethodGet, "/hello", h)

		actual, params := router.Inquiry(http.MethodTrace, "/hello")

		if actual != nil || len(params) != 0 {
			t.Errorf("\nhandler: expected nil, got %#v\nparams: expected {}, got %#v\n", actual, params)
		}
	})

	t.Run("should get nothing when inquiry for router given incorrect path", func(t *testing.T) {
		router := NewRouter()
		router.Register(http.MethodGet, "/hello", h)

		actual, params := router.Inquiry(http.MethodGet, "/hello/world")

		if actual != nil || len(params) != 0 {
			t.Errorf("\nhandler: expected nil, got %#v\nparams: expected {}, got %#v\n", actual, params)
		}
	})
}

//func TestDynamicRouter(t *testing.T) {}
