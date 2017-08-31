package client

import (
	"encoding/json"
)

type Request struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

func NewRequest(method string, params Params) *Request {
	pb, _ := json.Marshal(params)
	return &Request{
		Method: method,
		Params: pb,
	}
}

type Params []interface{}
