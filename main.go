package main

import "fmt"

func main() {

	foo(1, 2, 3, 4, 5)
	r := sum(1, 2, 3, 4, 5)
	fmt.Printf("total = %v \n", r)
}

func foo(x ...int) {
	// x, will be stored as a slice of ints, []int
	fmt.Println(x)
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
