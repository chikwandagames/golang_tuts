package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("anonymous func 1")
	}()

	func(x int) {
		fmt.Println("anonymous func 1:", x)
	}(20)
}
