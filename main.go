package main

import (
	"fmt"
)

// untyped constants are known as constant of a kind
const a = 42
const b = 42.47
const c = "James Bond"

// Alternaive definistion
const (
	d        = 34
	e string = "Wisdom" // this const is typed, therefore cannot be changed
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
