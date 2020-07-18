package main

import (
	"github.com/ltoddy/rabbit"
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
)

func main() {
	r := rabbit.NewRabbit(":2333")

	r.Get("/greet/<name>", func(request *request.Request) response.Response {
		return response.JsonResponse(rabbit.J{"name": request.Params.Get("name")})
	})

	r.Run()
}
