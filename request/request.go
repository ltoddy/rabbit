package request

import (
	"net/http"
)

type Request struct {
	*http.Request
	Params Params
}

func NewRequest(r *http.Request, params Params) *Request {
	return &Request{
		Request: r,
		Params:  params,
	}
}

type Params map[string]string

func (p Params) Get(field string) string {
	if value, ok := p[field]; ok {
		return value
	}
	return ""
}
