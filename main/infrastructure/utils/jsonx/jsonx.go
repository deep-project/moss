package jsonx

import (
	"encoding/json"
	"github.com/bytedance/sonic"
)

func Marshal(val interface{}) ([]byte, error) {
	return sonic.Marshal(val)
}

func Unmarshal(buf []byte, val interface{}) error {
	return sonic.Unmarshal(buf, val)
}

func MarshalString(val interface{}) (_ string, err error) {
	b, err := json.Marshal(val)
	if err != nil {
		return
	}
	return string(b), nil
}

func UnmarshalString(str string, val interface{}) error {
	return json.Unmarshal([]byte(str), val)
}
