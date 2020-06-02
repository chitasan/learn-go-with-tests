package main

import "testing"

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in French", func(t *testing.T) {
        got := Hello("Elodie", "French")
        want := "Bonjour, Elodie"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Italian", func(t *testing.T) {
        got := Hello("Elodie", "Italian")
        want := "Ciao, Elodie"
        assertCorrectMessage(t, got, want)
    })
}


// test file needs to be name_test.go
// test function must start with "Test" and function must be capitalized
// test function only takes one argument "t. *testing.T"
// t.Errorf pronts out a message and fail the test. 'f' stands for format, which allows ability to build a string with values inserted into the placeholder values %q
// t.Run is a subtest to describe diff scenarios 
