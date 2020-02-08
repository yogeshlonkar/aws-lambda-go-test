package main

import (
	"encoding/json"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	alt "github.com/yogeshlonkar/aws-lambda-go-test/local"
)

func TestSuccessIT(t *testing.T) {
	expected := "some-string"

	response, err := alt.Run(alt.Input{})

	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	var actual string
	if err = json.Unmarshal(response, &actual); err != nil {
		log.Println(err)
		t.FailNow()
	}

	if actual != expected {
		log.Println(err)
		t.FailNow()
	}
}

func TestSuccessITWithCustomPort(t *testing.T) {
	expected := "some-string"

	response, err := alt.Run(alt.Input{
		Port: 8818,
	})

	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	var actual string
	if err = json.Unmarshal(response, &actual); err != nil {
		log.Println(err)
		t.FailNow()
	}

	if actual != expected {
		log.Println(err)
		t.FailNow()
	}
}

func TestSuccessITWithCustomPortAndPath(t *testing.T) {
	expected := "random-func-string"

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	response, err := alt.Run(alt.Input{
		Port:          9922,
		AbsLambdaPath: path.Join(basepath, "test/random_name_main.go"),
	})

	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	var actual string
	if err = json.Unmarshal(response, &actual); err != nil {
		log.Println(err)
		t.FailNow()
	}

	if actual != expected {
		log.Println(err)
		t.FailNow()
	}
}
