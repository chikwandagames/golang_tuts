package main

import "fmt"

/*
	type error interface {
		Error() string
	}
*/

// Any type that implements Error() string is also of type error

func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	// Prints 6 (number of bytes) because of the new line character on printLN
	fmt.Println(n)
}
