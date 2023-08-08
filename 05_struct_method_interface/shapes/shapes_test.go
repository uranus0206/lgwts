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
	aeraTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.16},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range aeraTests {
		got := tt.shape.Aera()

		tolerance := 0.01
		if diff := math.Abs(got - tt.want); diff > tolerance {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
