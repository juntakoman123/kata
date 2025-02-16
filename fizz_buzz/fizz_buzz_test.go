package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("3の倍数はFizz", func(t *testing.T) {

		got := FizzBuzz(3)
		want := "Fizz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("5の倍数はBuzz", func(t *testing.T) {

		got := FizzBuzz(5)
		want := "Buzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("3と5の倍数はFizzBuzz", func(t *testing.T) {

		got := FizzBuzz(15)
		want := "FizzBuzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("3と5の倍数以外は元の数値", func(t *testing.T) {

		got := FizzBuzz(4)
		want := "4"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
