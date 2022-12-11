package main

import (
	"fmt"
	"math"
)

type square struct{
	// side is the length of a side of the square.
	side float64
}

// area returns the area of the square.
func (s square) area() float64 {
	return s.side * s.side
}

type circle struct{
	// radius is the radius of the circle.
	radius float64
}

// area returns the area of the circle.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// A shape is any type that has an area method.
type shape interface {
	area() float64
}

// info will print the area of a shape
func info(s shape) {
	fmt.Println("Area: ", s.area())
}

func main() {
	s := square{side: 6.66}
	info(s)

	c := circle{radius: 6.66}
	info(c)
}
