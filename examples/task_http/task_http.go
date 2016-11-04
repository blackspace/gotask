package task_http

import (
	"net/http"
	. "github.com/blackspace/gotask"
)

type HttpRequest struct {
	*TaskBase
	Request *http.Request
}

func (r *HttpRequest)Exec() interface{} {
	return "Hello World"
}

func NewHttpRequest() *HttpRequest {
	return &HttpRequest{TaskBase:NewTaskBase()}
}


