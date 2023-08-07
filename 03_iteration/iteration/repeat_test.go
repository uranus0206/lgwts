package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertMsg := func(t *testing.T, repeated, expected string) {
		t.Helper()

		if repeated != expected {
			t.Errorf("expected '%q', but got '%q'", expected, repeated)
		}
	}

	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertMsg(t, repeated, expected)
	})
	t.Run("repeat 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertMsg(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5000)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
