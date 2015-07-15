package main

import (
	"fmt"
	"math"
)

// Shape is the interface that defines the behaviour of all shapes.
type Shape interface {
	area() float64
	perimeter() float64
}

// Coords represents a tuple of coordinates
type Coords struct {
	x, y float64
}

// Circle is the basic struct representing a circle.
type Circle struct {
	c Coords
	r float64
}

// Satisfies the Shape interface
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Satisfies the Shape interface
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

// Rectangle is the basic struct representing a rectangle.
type Rectangle struct {
	c    Coords
	w, h float64
}

// Satisfies the Shape interface
func (r *Rectangle) area() float64 {
	return r.w * r.h
}

// Satisfies the Shape interface
func (r *Rectangle) perimeter() float64 {
	return r.w*2 + r.h*2
}

// Calculates the total area of all the shapes given in parameter
func totalArea(shapes ...Shape) (area float64) {
	for _, s := range shapes {
		area += s.area()
	}
	return
}

func main() {
	c1 := Circle{c: Coords{x: 1, y: 1}, r: 1.5}
	fmt.Println("Circle area :", c1.area())
	r1 := Rectangle{c: Coords{x: 2, y: 2}, w: 1.3, h: 1}
	fmt.Println("Rectangle area :", r1.area())
	fmt.Println("Total area :", totalArea(&c1, &r1))
}
