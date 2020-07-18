package response

type rawResponse struct {
	code   int
	header map[string]string
	body   []byte
}

func newRawResponse(code int, header map[string]string, body []byte) *rawResponse {
	return &rawResponse{code, header, body}
}

func (r rawResponse) StatusCode() int {
	return r.code
}

func (r rawResponse) Header() map[string]string {
	return r.header
}

func (r rawResponse) Body() []byte {
	return r.body
}
