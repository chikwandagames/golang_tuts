package main

import "fmt"

//
func main() {
	// Composite literal = type{ values }
	// A composit literal, is an expression that creates a new instance
	// each time its evaluated,
	// they can be user to create Arrays, Structs, Slices and maps
	x := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(x)

	x = append(x, 8, 9, 10)
	fmt.Println(x)

	y := []int{11, 12, 13, 14, 15}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("deleting 4 and position 3 in the slice")
	x = append(x[:4], x[5:]...)
	fmt.Println(x)

}
