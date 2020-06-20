package rabbit

import (
	"encoding/json"
	"fmt"
	"github.com/ltoddy/rabbit/types"
	"net/http"
)

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	ctx := new(Context)
	ctx.writer = writer
	ctx.request = request
	return ctx
}

func (c *Context) Json(v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := c.writer.Write(data); err != nil {
		panic(err)
	}
}

func (c *Context) Html(content string) {
	c.writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := fmt.Fprint(c.writer, content); err != nil {
		panic(err)
	}
}

func (c *Context) Text(content string) {
	c.writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := fmt.Fprint(c.writer, content); err != nil {
		panic(err)
	}
}

func (c *Context) Redirect(p types.Path) {}
