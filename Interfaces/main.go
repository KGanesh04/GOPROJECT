package main

import "fmt"

// Define an Interface
type Shape interface {
	Area() float64
}

// Define a Struct (Rectangle)
type Rectangle struct {
	Width, Height float64
}

// Implement Area() Method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Define Another Struct (Circle)
type Circle struct {
	Radius float64
}

// Implement Area() Method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	//  Create Interface Type Variable
	var s Shape

	// Assign Rectangle to Interface
	s = Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle Area:", s.Area())

	// Assign Circle to Interface
	s = Circle{Radius: 7}
	fmt.Println("Circle Area:", s.Area())
}
