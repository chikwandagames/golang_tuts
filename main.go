package main

import (
	"fmt"
	"math"
)

// POLYMORPHSM, the area() method behave differently depending
// on the type of shape you call it using

// Square and Circle are concreate types
// Square is ...
type Square struct {
	side float64
}

// Circle is ...
type Circle struct {
	radius float64
}

// Shape is an INTERFACE, abstract type
// so now anything that has the area() float64 method
// implements the Shape interface
type Shape interface {
	area() float64
}

// Here we attach a method to the type Square
// i.e. Square now implements the shape interface
func (z Square) area() float64 {
	return z.side * z.side
}

// Here we attach a method to the type circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// info() can take an circle or square
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)
	c := Circle{5}
	info(c)
}
