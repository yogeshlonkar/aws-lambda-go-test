package util

import "testing"

type TypeForTest struct {
	Something int
}

func TestTryUnmarshal(t *testing.T) {
	actual := TypeForTest{}
	TryUnmarshal([]byte(`{ "something": 1234 }`), &actual)
	if actual.Something != 1234 {
		t.Errorf("actual.Something=%d, want 1234", actual.Something)
	}
}

func TestTryUnmarshalEscaped(t *testing.T) {
	actual := TypeForTest{}
	TryUnmarshalEscaped([]byte(`"{ \"something\": 1234 }"`), &actual)
	if actual.Something != 1234 {
		t.Errorf("actual.Something=%d, want 1234", actual.Something)
	}
}
