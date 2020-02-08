package main

import (
	"encoding/json"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func TestSuccessIT(t *testing.T) {
	expected := "some-string"

	response, err := Run(Input{})

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

	response, err := Run(Input{
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

	response, err := Run(Input{
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
