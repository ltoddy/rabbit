package rabbit

import "github.com/ltoddy/rabbit/handler"

type Any = interface{}

type J = map[string]Any

type entryHandler struct {
	fullpath string
	fn       handler.HandlerFunction
}
