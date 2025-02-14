package response

import "net/http"

//This is the response struct for all the responses
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// New create a new response
func New(status int, message string, data any, error any) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   error,
	}
}

func BindJSONErr(err error) Response {
	return New(http.StatusBadRequest, "failed to bind json", nil, err.Error())
}

func BindQueryErr(err error) Response {
	return New(http.StatusBadRequest, "failed to bind query", nil, err.Error())
}
func BindPathParamErr(err error) Response {
	return New(http.StatusBadRequest, "failed to bind path param", nil, err.Error())
}
