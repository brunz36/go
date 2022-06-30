package main

import "fmt"

func main() {
	s := square{6}
	printArea(s)

	t := triangle{10, 6}
	printArea(t)
}

// An interface for all types of shape
type shape interface {
	getArea() float64
}

// Print the area of a give shape using the shape.getArea method
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// An implementation of shape with 4 equal sides
type square struct {
	sideLength float64
}

// An implementation of shape.getArea to get the area of a square
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// An implementation of shape with 3 sides
type triangle struct {
	height, base float64
}

// An implementation of shape.getArea to get the area of a triangle
func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

