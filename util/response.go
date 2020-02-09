package util

import (
	"encoding/json"
)

type Response struct {
	StatusCode        int
	Headers           interface{}
	MultiValueHeaders interface{}
	Body              interface{}
}

func (res *Response) PaseBody(bodyStruct interface{}) {
	bodyTxt := res.Body.(*json.RawMessage)
	TryUnmarshalEscaped(*bodyTxt, bodyStruct)
}

func (res *Response) PaseHeaders(headStruct interface{}) {
	headerTxt := res.Headers.(*json.RawMessage)
	TryUnmarshal(*headerTxt, headStruct)
}

func ParseResponse(responseTxt []byte) *Response {
	var headers json.RawMessage
	var body json.RawMessage
	response := Response{
		Headers: &headers,
		Body:    &body,
	}
	TryUnmarshal(responseTxt, &response)
	return &response

}
