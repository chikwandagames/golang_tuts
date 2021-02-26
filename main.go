package main

import "fmt"

// CALLBACKS
// Passing a function as an argument

func main() {

	ii := []int{1, 2, 3, 4, 5}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers sum:", s2)

}

func sum(xi ...int) int {

	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {

			yi = append(yi, v)
		}
	}
	fmt.Println("even numbers:", yi)
	// f is sum() in this case
	return f(yi...)
}

// func filterEvens(xi ...int) []int {

// 	ls := []int{}

// 	for _, v := range xi {
// 		if v%2 == 0 {
// 			ls = append(ls, v)
// 		}
// 	}
// 	return ls
// }

// func addEvens(f func(xi ...int) []int, vi ...int) int {
// 	vl := f(vi...)
// 	r := sum(vl...)
// 	return r
// }
