package main

import (
	"github.com/ltoddy/rabbit"
	"github.com/ltoddy/rabbit/response"
	"net/http"
)

func main() {
	r := rabbit.NewRabbit(":2333")

	r.Get("/", func(request *http.Request) response.Response {
		return response.JsonResponse(rabbit.J{
			"hello": "world",
		})
	})

	r.Get("/hello", func(request *http.Request) response.Response {
		return response.TextResponse("Hello world")
	})

	r.Run()
}
