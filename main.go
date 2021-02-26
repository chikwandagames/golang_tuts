package main

import "fmt"

//
func main() {
	// Composite literal = type{ values }
	// A composit literal, is an expression that creates a new instance
	// each time its evaluated,
	// they can be user to create Arrays, Structs, Slices and maps
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)

	// up to but not including
	fmt.Println(x[:3])

	// from position 2 to the end
	fmt.Println(x[2:])

	// from start to second last element
	fmt.Println(x[:len(x)-1])

}
