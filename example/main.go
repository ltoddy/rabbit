package main

import (
	"github.com/ltoddy/rabbit"
	"net/http"
)

func main() {
	r := rabbit.NewRabbit(":2333")

	r.Get("/", func(request *http.Request) rabbit.Response {
		return rabbit.JsonResponse(rabbit.J{
			"hello": "world",
		})
	})

	r.Get("/hello", func(request *http.Request) rabbit.Response {
		return rabbit.TextResponse("Hello world")
	})

	r.Run()
}
