package structs // Variável receptora
import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

/*
	É uma convenção em Go que a variável receptora seja a primeira letra do tipo.
	https://go.dev/ref/spec#Identifiers
	https://google.github.io/styleguide/go/decisions#receiver-names
	https://google.github.io/styleguide/go/decisions.html#single-letter-variable-names
*/

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return math.Pi * (2 * c.Radius)
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
