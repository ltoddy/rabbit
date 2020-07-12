package response

type jsonResponse struct {
	data []byte
	code int // http status code
}

func NewJsonResponse(data []byte, code int) *jsonResponse {
	return &jsonResponse{data: data, code: code}
}

func (j *jsonResponse) StatusCode() int {
	return j.code
}

func (j *jsonResponse) Header() map[string]string {
	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=utf-8"
	header["Server"] = "github.com/ltoddy/rabbit"
	return header
}

func (j *jsonResponse) Body() []byte {
	return j.data
}
