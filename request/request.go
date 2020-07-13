package request

import (
	"net/http"
)

type Request struct {
	*http.Request
	Params Params
}

type Params map[string]string
