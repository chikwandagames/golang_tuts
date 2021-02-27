package main

import "fmt"

//
func main() {
	a := 34
	// The address of "a" is a pointer to an int
	fmt.Printf("\"a\" memory address: %v\n", &a)
	fmt.Printf("\"a\" type: %T \n", &a)

	// b points to the address of a
	var b *int = &a
	// c points to the address of a
	c := &a
	fmt.Println(b)
	fmt.Println(c)

	// Dereferencing an address
	fmt.Println(*c)
	fmt.Println(*b)

	fmt.Println(*&a)
	fmt.Println(*&b)

	// Dereference b and assign new value
	*b = 100
	fmt.Println(a)

}
