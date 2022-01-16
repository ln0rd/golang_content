package forms

import (
	"fmt"
	"math"
)

type Form interface {
	Area() float64
}

func writeArea(f Form)  {
	fmt.Printf("Area form is %0.2f \n", f.Area())
}

type Rectangle struct {
	Height float64
	Width float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func main()  {
	r := Rectangle{ Height: 10, Width: 15 }
	writeArea(r)

	c := Circle{ Radius: 100 }
	writeArea(c)
}