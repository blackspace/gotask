package dohttp

import (
	. "github.com/blackspace/gotask"
	"net/http"
)



type DoRequest struct {
	*TaskBase
	Request *http.Request
}

func (r *DoRequest)Exec() interface{} {
	return "Hello World"
}

func NewDoRequest() *DoRequest {
	return &DoRequest{TaskBase:NewTaskBase()}
}


