package error

import "net/http"

type Err struct {
	Code int
	Msg  string
}

func (e *Err) Error() string {
	return e.Msg
}

var statusCode = map[int]int{
	1001: http.StatusBadGateway,
}

func HttpStatusCode(code int) int {
	v, ok := statusCode[code]
	if ok {
		return v
	}
	return http.StatusOK
}
