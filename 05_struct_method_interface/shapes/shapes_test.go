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
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.16},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range aeraTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Aera()

			tolerance := 0.01
			if diff := math.Abs(got - tt.want); diff > tolerance {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})

	}
}
