package handler

import (
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
)

type Handler interface {
	Serve(*request.Request) response.Response
}

type HandlerFunction func(*request.Request) response.Response

func (f HandlerFunction) Serve(r *request.Request) response.Response {
	return f(r)
}
