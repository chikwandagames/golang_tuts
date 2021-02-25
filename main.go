package main

import (
	"fmt"
)

// iota is good if you need a number that auto increments
const (
	a = iota
	b = iota
	c = iota
)

// Alternatively
// this resets
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
