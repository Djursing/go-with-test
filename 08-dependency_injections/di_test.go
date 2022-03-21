package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	exp := "Hello, Chris"

	if got != exp {
		t.Errorf("\ngot: %q\nexp: %q", got, exp)
	}
}
