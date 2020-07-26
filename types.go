package rabbit

type Any = interface{}

type J = map[string]Any

type entryHandler struct {
	fullpath string
	fn       HandlerFunction
}
