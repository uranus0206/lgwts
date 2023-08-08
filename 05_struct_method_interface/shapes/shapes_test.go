package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAera(t *testing.T) {

	checkAera := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Aera()

		tolerance := 0.001
		if diff := math.Abs(got - want); diff > tolerance {
			t.Errorf("got %f want %f diff %f", got, want, diff)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  12.0,
			Height: 6.0,
		}
		checkAera(t, rectangle, 72.0)

	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10}
		checkAera(t, circle, 314.16)
	})
}
