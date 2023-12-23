package structs

import "math"

// Shape is implemented by anything that can tell us its Area.
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the perimeter of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle reoresents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents the dimentions of a triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
