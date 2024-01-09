package test

import (
	"tdd-go/pkg"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Say hello to Omar", func(t *testing.T) {
		want := "Hello, Omar!"
		got := pkg.Hello("Omar", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("When an empty string is provided, say 'Hello, World!'", func(t *testing.T) {
		want := "Hello, World!"
		got := pkg.Hello("", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		want := "Hola, Omar!"
		got := pkg.Hello("Omar", "Spanish")

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		want := "Bonjour, Omar!"
		got := pkg.Hello("Omar", "French")

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
