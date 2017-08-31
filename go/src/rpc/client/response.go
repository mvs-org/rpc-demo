package client

import (
	"encoding/json"
	"errors"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
	Result  string `json:"result"`
}

func (ex *ResponseError) UnmarshalError(body []byte) error {
	json.Unmarshal(body, ex)
	if ex.Code == 0 {
		return errors.New("exception was not caught")
	}
	return nil
}
