package main

import "fmt"

// CLOSURE
// Closure is about enlosing the scope of a variable to make sure
// its scope is limited

/*
	A closure is function value that references variables from outside its own
  function body,
	i.e. and inner function that remembers and has access to variables in the local
	scope in which it was created even after the outer function is done executin
*/

// increment() returns a fuction that returns an int
func increment() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := increment()
	// The value keep incrementing because the inner anonymous function is
	// keeping track of the value i in its outer scope, even though increment()
	// is done executing
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	print("\n")

	nextInt2 := increment()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
}
