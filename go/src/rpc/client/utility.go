package client

import (
	"bytes"
	"encoding/json"
	"github.com/antonholmquist/jason"
)

func childUnmarshal(data []byte, result interface{}, key ...string) error {
	object, err := jason.NewObjectFromBytes(data)
	if err != nil {
		return err
	}
	child, err := object.GetObject(key...)
	if err != nil {
		return err
	}
	childData, err := child.Marshal()
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader([]byte(childData))).Decode(&result)
}
