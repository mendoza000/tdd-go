package test

import (
	"tdd-go/pkg"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, Omar!"
	got := pkg.Hello("Omar")

	if got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
