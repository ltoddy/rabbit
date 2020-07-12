package response

type textResponse struct {
	content string
	code    int // http status code
}

func NewTextResponse(content string, code int) *textResponse {
	return &textResponse{content: content, code: code}
}

func (t *textResponse) StatusCode() int {
	return t.code
}

func (t *textResponse) Header() map[string]string {
	header := make(map[string]string)
	header["Content-Type"] = "text/plain; charset=utf-8"
	header["Server"] = "github.com/ltoddy/rabbit"
	return header

}

func (t *textResponse) Body() []byte {
	return []byte(t.content)
}
