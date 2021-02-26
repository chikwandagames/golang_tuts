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

	for i, v := range x {
		fmt.Printf("%v: %v\n", i, v)
	}

}
