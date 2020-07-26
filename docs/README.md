# Rabbit

Rabbit is pure Golang implementation for web framework that's written to go fast.

The goal of the project is to provide a simple way to get up and running a highly
performant HTTP server that is easy to build, to expand, and ultimately to scale.

Rabbit is developed [on Github](http://github.com/ltoddy/rabbit). Contributions are welcome!

# Rabbit aspires to be simple

```go
package main

import (
	"github.com/ltoddy/rabbit"
	"github.com/ltoddy/rabbit/request"
	"github.com/ltoddy/rabbit/response"
)

func main() {
	r := rabbit.NewRabbit(":2333")
	r.Get("/", func(r *request.Request) response.Response {
		return response.TextResponse("Hello world!\n")
	})
	r.Run()
}
```

# Guides

- Getting Started
  1. Install Rabbit
  2. Create a file called `main.go`
  3. Run the server
  4. Check your browser

TODO
