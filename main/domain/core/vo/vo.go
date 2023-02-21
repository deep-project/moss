package vo

import (
	"database/sql/driver"
	"encoding/json"
)

// StringArray 字段类型
type StringArray []string

func (val *StringArray) Scan(value interface{}) error {
	s, _ := value.(string)
	if s == "" {
		*val = []string{}
		return nil
	}
	_ = json.Unmarshal([]byte(s), val)
	return nil
}

func (val StringArray) Value() (_ driver.Value, err error) {
	b, err := json.Marshal(val)
	if err != nil {
		return
	}
	return string(b), err
}
