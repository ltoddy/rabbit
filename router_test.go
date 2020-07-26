package rabbit

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
		router := newRouter()
		router.register(http.MethodGet, "/hello/world", h)

		actual, params := router.inquiry(http.MethodGet, "/hello/world")

		if actual != h {
			t.Errorf("handler: expected handler, got %#v\n", actual)
		}

		if len(params) != 0 {
			t.Errorf("params: expected {}, got %#v\n", params)
		}
	})

	t.Run("should get nothing when inquiry for router given incorrect http method", func(t *testing.T) {
		router := newRouter()
		router.register(http.MethodGet, "/hello", h)

		actual, params := router.inquiry(http.MethodTrace, "/hello")

		if actual != nil || len(params) != 0 {
			t.Errorf("\nhandler: expected nil, got %#v\nparams: expected {}, got %#v\n", actual, params)
		}
	})

	t.Run("should get nothing when inquiry for router given incorrect path", func(t *testing.T) {
		router := newRouter()
		router.register(http.MethodGet, "/hello", h)

		actual, params := router.inquiry(http.MethodGet, "/hello/world")

		if actual != nil || len(params) != 0 {
			t.Errorf("\nhandler: expected nil, got %#v\nparams: expected {}, got %#v\n", actual, params)
		}
	})
}

func TestDynamicRouter(t *testing.T) {
	h1 := handle{}
	h2 := handle{}

	t.Run("", func(t *testing.T) {
		router := newRouter()
		router.register(http.MethodGet, "/", h1)

		actualHandler, _ := router.inquiry(http.MethodGet, "/")

		if actualHandler != h1 {
			t.Errorf("expected: %#v, got: %#v\n", h1, actualHandler)
		}
	})

	t.Run("", func(t *testing.T) {
		router := newRouter()
		router.register(http.MethodGet, "/hello/world", h1)
		router.register(http.MethodGet, "/hello/world/<name>", h2)
		expectedParams := request.Params{"name": "ltoddy"}

		actualHandler, actualParams := router.inquiry(http.MethodGet, "/hello/world/ltoddy")

		if len(actualParams) != len(expectedParams) && expectedParams["name"] != actualParams["name"] {
			t.Errorf("expected: %#v, got: %#v\n", expectedParams, actualParams)
		}

		if actualHandler != h2 {
			t.Errorf("expected: %#v, got: %#v\n", h2, actualHandler)
		}
	})

	t.Run("", func(t *testing.T) {
		router := newRouter()
		router.register(http.MethodGet, "/<id>/<name>", h1)
		expectedParams := request.Params{"id": "10", "name": "ltoddy"}

		actualHandler, actualParams := router.inquiry(http.MethodGet, "/10/ltoddy")

		if actualHandler != h1 {
			t.Errorf("expected: %#v, got: %#v\n", h1, actualHandler)
		}

		if len(actualParams) != 2 || actualParams["id"] != expectedParams["id"] || actualParams["name"] != expectedParams["name"] {
			t.Errorf("expected: %#v, got: %#v\n", expectedParams, actualParams)

		}
	})
}
