package ijson

import (
	"encoding/json"
)

var (
	emptyBytes = []byte("{}")
)

func ToJsonByte(message interface{}) []byte {
	if data, err := json.Marshal(message); err == nil {
		return data
	}
	return emptyBytes
}

func ToJsonString(message interface{}) string {
	return string(ToJsonByte(message))
}
