package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

func writeArea(f form)  {
	fmt.Printf("Area form is %0.2f \n", f.area())
}

type rectangle struct {
	height float64
	width float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main()  {
	r := rectangle{ height: 10, width: 15 }
	writeArea(r)

	c := circle{ radius: 100 }
	writeArea(c)
}