package structsmethodsinterfaces

import "math"

// ===================
// General Shape Interface
// ===================
type Shape interface {
	Area() float64
}

// ===================
// RECTANGLE + METHODS
// ===================
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// ===================
// CIRCLE + METHODS
// ===================
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ===================
// TRIANGLE + METHODS
// ===================
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
