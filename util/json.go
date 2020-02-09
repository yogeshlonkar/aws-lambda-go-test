package util

import (
	"encoding/json"
	"strconv"
)

func TryUnmarshal(text []byte, response interface{}) {
	err := json.Unmarshal(text, response)
	if err != nil {
		panic(err)
	}
}

func TryUnmarshalEscaped(jsonStr json.RawMessage, response interface{}) {
	sB, _ := strconv.Unquote(string(jsonStr))
	err := json.Unmarshal([]byte(sB), response)
	if err != nil {
		panic(err)
	}
}
