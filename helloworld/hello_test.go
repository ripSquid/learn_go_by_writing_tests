package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to recipient", func(t *testing.T) {
		got := Hello("Peter", "English")
		want := "Hello, Peter"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Griffin", "Spanish")
		want := "Hola, Griffin"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Peter", "French")
		want := "Bonjour, Peter"
		assertCorrectMessage(t, got, want)
	})
}

func TestGreetingPrefix(t *testing.T) {
	t.Run("English default prefix", func(t *testing.T) {
		got := greetingPrefix("")
		want := "Hello, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("Spanish Prefix", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola, "
		assertCorrectMessage(t, got, want)
	})

	t.Run("French Prefix", func(t *testing.T) {
		got := greetingPrefix("French")
		want := "Bonjour, "
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
