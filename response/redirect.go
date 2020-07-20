package response

type redirectResponse struct {
	path string
	code int
}

func newRedirectResponse(rawurl string, code int) *redirectResponse {
	// TODO: compatible relative path and absolute path
	return &redirectResponse{rawurl, code}
}

func (r *redirectResponse) StatusCode() int {
	return r.code
}

func (r *redirectResponse) Header() map[string]string {
	header := make(map[string]string)
	header["Location"] = r.path
	header["Content-Type"] = "text/html; charset=utf-8"
	return header
}

func (r *redirectResponse) Body() []byte {
	return []byte{}
}
