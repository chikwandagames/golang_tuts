package main

import "fmt"

// var lets you decalare a global variable
var n = 45

// declares a variable z of type int, with a default ZERO value
var z int
var s = "Shaken, not stirred"

// String literal
var m = `Hey player`

func main() {
	x := 42

	fmt.Println(x)
	// Print variable type
	fmt.Printf("%T \n", x)

	fmt.Println(s)
	// Print variable type
	fmt.Printf("%T \n", s)

	fmt.Println(m)

}
