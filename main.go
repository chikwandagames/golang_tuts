package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7}

	// Slice DESTRUCTURING
	x := sum(xi...)
	fmt.Println(x)
	fmt.Println("")

	foo(xi...)
	foo(6, 7, 8, 9)
	foo()

}

// You can pass zero or more ints to this function
// The final parameter has to be 0 or more
// foo(s string, x ...int)  GOOD
// foo(x ...int, s string)  BAD
func foo(x ...int) {
	// x, will be stored as a slice of ints, []int
	fmt.Println(x)
	fmt.Printf("length: %v \n", len(x))
	fmt.Printf("%T\n", x)
}

// x, will be stored as a slice of ints, []int
func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		fmt.Printf("%v: %v + %v = %v \n", i, sum, v, v+sum)
		sum += v
	}
	return sum
}
