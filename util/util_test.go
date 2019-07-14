package util

import "testing"

/*
// function not used.
func TestStringToDate(t *testing.T) time.Time{
	assertCorrectResponse := func(t *testing.T, got, want ) {
		t.Helper()
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("parsing a number", func(t *testing.T) {
		got := BetterAtoi("5")
		want := 5
		assertCorrectResponse(t, got, want)
	})

	t.Run("parsing non-numeric values", func(t *testing.T) {
		got := BetterAtoi("~")
		want := ^int(0)
		assertCorrectResponse(t, got, want)
	})

	t.Run("parsing negative numbers", func(t *testing.T) {
		got := BetterAtoi("-5")
		want := -5
		assertCorrectResponse(t, got, want)
	})

}*/

func TestBetterAtoi(t *testing.T) {
	assertCorrectResponse := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("parsing a number", func(t *testing.T) {
		got := BetterAtoi("5")
		want := 5
		assertCorrectResponse(t, got, want)
	})

	t.Run("parsing non-numeric values", func(t *testing.T) {
		got := BetterAtoi("~")
		want := ^int(0)
		assertCorrectResponse(t, got, want)
	})

	t.Run("parsing negative numbers", func(t *testing.T) {
		got := BetterAtoi("-5")
		want := -5
		assertCorrectResponse(t, got, want)
	})

}
