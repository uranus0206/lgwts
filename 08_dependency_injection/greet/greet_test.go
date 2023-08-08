package greet

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("write to bytes buffer", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	t.Run("write to stdout", func(t *testing.T) {
		Greet(os.Stdout, "Chris")
	})
}
