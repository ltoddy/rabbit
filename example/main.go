package main

import (
	"github.com/ltoddy/rabbit"
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
)

func main() {
	r := rabbit.NewRabbit(":2333")

	r.Get("/", func(request *request.Request) response.Response {
		return response.JsonResponse(rabbit.J{
			"hello": "world",
		})
	})

	r.Get("/hello", func(request *request.Request) response.Response {
		return response.TextResponse("Hello world")
	})

	r.Run()
}
