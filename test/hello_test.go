package test

import (
	"tdd-go/pkg"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Say hello to Omar", func(t *testing.T) {
		want := "Hello, Omar!"
		got := pkg.Hello("Omar")

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("When an empty string is provided, say 'Hello, World!'", func(t *testing.T) {
		want := "Hello, World!"
		got := pkg.Hello("")

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
