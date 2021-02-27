package main

import "fmt"

// When to use pointer
// 1.
// Pointers a good when don't want to pass large pieces of data
// pointers let you simpley pass the address
// this is good for performance
// 2.
// When you need to modify something at a certain location

func main() {
	a := 0
	foo(a)
	fmt.Println(a)

	fmt.Println("")

	fooPtr(&a)
	fmt.Println(a)

}

func foo(x int) {
	fmt.Println(x)
	x = 43
	fmt.Println(x)
}

func fooPtr(x *int) {
	fmt.Println(x)
	// Set the value an &x to 100
	*x = 100
	fmt.Println(x)
}
