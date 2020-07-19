package response

import (
	"encoding/json"
	"net/http"
)

// low level struct
type Response interface {
	StatusCode() int

	/**
	response header fields:
	e.g.:
	- Date: Mon, 27 Jul 2009 12:28:53 GMT
	- Server: Apache/2.2.14 (Win32)
	- Last-Modified: Wed, 22 Jul 2009 19:15:56 GMT
	- Content-Length: 88
	- Content-Type: text/html
	- Connection: Closed
	*/
	Header() map[string]string

	/**
	optional a message body
	*/
	Body() []byte
}

func JsonResponse(v interface{}) Response {
	data, err := json.Marshal(v)
	if err != nil {
		return newJsonResponse([]byte(err.Error()), http.StatusInternalServerError)
	}

	return newJsonResponse(data, http.StatusOK)
}

func TextResponse(content string) Response {
	return newTextResponse(content, 200)
}

func RawResponse(code int, header map[string]string, body []byte) Response {
	return newRawResponse(code, header, body)
}

func Redirect(path string) Response {
	//TODO
	panic("unimplemented!")
}

func Fail(code int) Response {
	//TODO
	panic("unimplemented!")
}

func Render(tmpl string, args interface{}) Response {
	//TODO
	panic("unimplemented!")
}
