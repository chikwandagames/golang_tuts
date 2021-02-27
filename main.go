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

	// Composite literal = type{ values }
	// A composit literal, is an expression that creates a new instance
	// each time its evaluated,
	// they can be user to create Arrays, Structs, Slices and maps
	ii := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ii)

	// make(), lets you specify, (type, length, capacity)
	// because the underlining data type for a sclice is an array,
	// capacity creates an array with x position, for optimisation
	// by so doing the compiler doest have to keep recreating arrays
	// each time the length gets bigger than the capacity
	i2 := make([]int, 10, 15)
	fmt.Println(i2)
	fmt.Println(len(i2))
	fmt.Println(cap(i2))

	// Adding an item will change the length but not the capacity
	i2 = append(i2, 23)
	fmt.Println(i2)
	fmt.Println(len(i2))
	fmt.Println(cap(i2))

	// If we append items to the array and the length out grows the capacity
	// the capacity doubles
	i2 = append(i2, 23)
	i2 = append(i2, 23)
	i2 = append(i2, 23)
	i2 = append(i2, 23)
	i2 = append(i2, 23)
	fmt.Println(i2)
	fmt.Println(len(i2))
	fmt.Println(cap(i2))

}
