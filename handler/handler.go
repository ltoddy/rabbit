package handler

import (
	"github.com/ltoddy/rabbit/response"
	"net/http"
)

type Handler interface {
	Serve(*http.Request) response.Response
}

type HandlerFunction func(*http.Request) response.Response

func (f HandlerFunction) Serve(r *http.Request) response.Response {
	return f(r)
}
