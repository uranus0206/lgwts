package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %d", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v want %v", got, want)
	// }

	if len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	} else {
		for i, v := range want {
			if v != got[i] {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}
}
