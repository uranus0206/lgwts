package shapes

import (
	"math"
)

type Shape interface {
	Aera() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Aera() (aera float64) {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Aera() (aera float64) {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rect Rectangle) (peri float64) {
	return 2 * (rect.Height + rect.Height)
}

func Aera(rect Rectangle) (aera float64) {
	return (rect.Height * rect.Width)
}
