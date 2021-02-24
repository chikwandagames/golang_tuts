package main

import "fmt"

var a int

// Custom type int
type hotdog int

var b hotdog

func main() {

	// Type conversion, coverting "b" into an int, which is its
	// underlying type
	c := int(b)

	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
}
