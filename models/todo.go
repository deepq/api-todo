package models

import (
	"io"

	"github.com/google/uuid"
)

type Todo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"title"`
	Done bool      `json:"completed"`
}

type ServiceAction func(ApiRequest, *ServiceResponse) error
type ServiceResponse struct {
	Response interface{}
}

type ApiRequest struct {
	Vars  map[string]string
	Query map[string][]string
	Body  io.ReadCloser
}

type ApiResponse struct {
	Err   string      `json:"error"`
	Data  interface{} `json:"data"`
	ReqID interface{} `json:"reqID"`
}
