package util

import "testing"

func TestParseResponse(t *testing.T) {
	actual := ParseResponse([]byte(`{
		"statusCode":200,
		"headers":null,
		"multiValueHeaders":null,
		"body":"{\"someArr\":[\"\"]}"
	}`))
	if actual.StatusCode != 200 {
		t.Errorf("actual.StatusCode=%d, want 200", actual.StatusCode)
	}
}

type TestHeaderStruct struct {
	ContentType string
}

type TestBody struct {
	SomeArr []string
}

func TestResponse_PaseHeaders(t *testing.T) {
	response := ParseResponse([]byte(`{
		"statusCode":200,
		"headers": {
		  "content-type": "application/json"
		},
		"multiValueHeaders":null,
		"body":"{\"someArr\":[\"\"]}"
	}`))
	actualHeaders := TestHeaderStruct{}
	response.PaseHeaders(&actualHeaders)
	if actualHeaders.ContentType == "application/json" {
		t.Errorf("actualHeaders.ContentType=%s, wanted application/json", actualHeaders.ContentType)
	}

	actualBody := TestBody{}
	response.PaseBody(&actualBody)
	if len(actualBody.SomeArr) != 1 {
		t.Errorf("len(actualBody.SomeArr)=%d, wanted 1", len(actualBody.SomeArr))
	}
}
