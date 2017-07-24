package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	addr string
	path string
}

func NewClient(addr, path string) *Client {
	return &Client{addr, addr + path}
}

func (c *Client) Request(method string, params Params) (json.RawMessage, *ResponseError) {
	pb, _ := json.Marshal(params)
	req := &Request{
		Method: method,
		Params: pb,
	}

	reqJson, _ := json.Marshal(req)
	fmt.Println(string(reqJson))
	httpreq, _ := http.NewRequest("POST", c.path, bytes.NewBuffer(reqJson))
	httpreq.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := http.DefaultClient.Do(httpreq)
	defer func(req *http.Request, resp *http.Response) {
		if req != nil && req.Body != nil {
			req.Body.Close()
		}
		resp.Body.Close()
	}(httpreq, resp)

	ex := new(ResponseError)
	if resp.StatusCode != http.StatusOK {
		ex.Message = "http status code: " + strconv.Itoa(resp.StatusCode)
		return nil, ex
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		ex.Message = "http response read faild"
		return nil, ex
	}
	if ex.UnmarshalError(body) != nil {
		return body, nil
	}
	return nil, ex
}
