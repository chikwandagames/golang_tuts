package main

import "fmt"

func main() {
	x := 1
	// for statemaent with single condition
	for x < 5 {
		fmt.Println(x)
		x++
	}

	fmt.Println("")
	// this for is like a while loop
	j := 1
	for {
		if j > 5 {
			break
		}
		fmt.Println(j)
		j++
	}

	fmt.Println("")
	// for statement with a clause
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
