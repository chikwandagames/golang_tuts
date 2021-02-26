package main

import (
	"fmt"
)

func main() {
	x := bar()
	fmt.Printf("%T \n", x)

	i := x()
	fmt.Println(i)

	// Alternatively
	// Here we execute both bar() and its returned function
	f := bar()()
	fmt.Println(f)

}

// bar(), returns a function that returns and int
// functions are types in Golang
func bar() func() int {
	return func() int {
		return 500
	}
}
