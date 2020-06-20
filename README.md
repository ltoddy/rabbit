# rabbit


### Quick Start

> go get github.com/ltoddy/rabbit 

*main.go*

```go
package main

import (
	"github.com/ltoddy/rabbit"
)

func main() {
	r := rabbit.NewRabbit(":2333")

	r.Get("/", func(ctx *rabbit.Context) {
		ctx.Text("Hello world!")
	})

	r.Run()
}
```

```
$ curl -i localhost:2333
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Sat, 20 Jun 2020 03:07:21 GMT
Content-Length: 12

Hello world!
```
