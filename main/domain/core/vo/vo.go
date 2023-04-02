package vo

import (
	"database/sql/driver"
	"encoding/json"
)

// StringArray 字段类型
type StringArray []string

func (val *StringArray) Scan(value interface{}) error {
	// mysql下，从数据库返回的是[]byte,而sqlite返回的是string
	b, ok := value.([]byte)
	if !ok {
		s, _ := value.(string)
		b = []byte(s)
	}
	if len(b) == 0 {
		*val = []string{}
		return nil
	}
	_ = json.Unmarshal(b, val)
	return nil
}

func (val StringArray) Value() (_ driver.Value, err error) {
	b, err := json.Marshal(val)
	if err != nil {
		return
	}
	return string(b), err
}
