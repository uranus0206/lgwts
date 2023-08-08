package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(x, y float64) (peri float64) {
	return 2 * (x + y)
}

func Aera(x, y float64) (aera float64) {
	return (x * y)
}
