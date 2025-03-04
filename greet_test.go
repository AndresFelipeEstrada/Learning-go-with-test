package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Andres")

	got := buffer.String()
	want := "Hello, Andres"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
