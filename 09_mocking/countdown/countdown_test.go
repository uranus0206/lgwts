package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("print 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &OperationCountdownSpy{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

	})

	t.Run("sleep after every print", func(t *testing.T) {
		spy := &OperationCountdownSpy{}
		Countdown(spy, spy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spy.Calls) {
			t.Errorf("want %v got %v", want, spy.Calls)
		}
	})
}
