package internal

import (
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
)

type Handler interface {
	Serve(*request.Request) response.Response
}
