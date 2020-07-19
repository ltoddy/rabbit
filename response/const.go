package response

type ResponseBehavior uint8

const (
	Success = ResponseBehavior(iota)
	Failure
)
