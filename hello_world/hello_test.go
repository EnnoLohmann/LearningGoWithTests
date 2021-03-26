package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("got %q but wanted %q", actual, expected)
		}
	}

	t.Run("should print Hello, world without params", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, world"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("should print Hello, Enno with Enno as param", func(t *testing.T) {
		actual := Hello("Enno", "")
		expected := "Hello, Enno"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("should print Hola, World with spanish as language", func(t *testing.T) {
		actual := Hello("", "Spanish")
		expected := "Hola, world"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("should print Bonjour, World with french as language", func(t *testing.T) {
		actual := Hello("", "French")
		expected := "Bonjour, world"

		assertCorrectMessage(t, actual, expected)
	})

}
