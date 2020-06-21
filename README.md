# rabbit


### Quick Start

> go get github.com/ltoddy/rabbit 

*main.go*

```go
package main

import (
	"fmt"
	"github.com/ltoddy/rabbit"
)

func main() {
	r := rabbit.NewRabbit(":2333")

	r.Get("/", func(ctx *rabbit.Context) {
		ctx.Text("Hello world!")
	})
	r.Get("/hello/:name", func(ctx *rabbit.Context) {
		ctx.Text(fmt.Sprintf("Hello %s\n", ctx.Params["name"]))
	})

	r.Run()
}
```

```
$ curl -i localhost:2333/hello/some
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Sun, 21 Jun 2020 08:51:28 GMT
Content-Length: 11

Hello some
```
