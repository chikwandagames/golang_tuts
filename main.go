package main

import "fmt"

var y = 43

func main() {

	fmt.Println(y)
	fmt.Printf("%v \n", y)
	fmt.Printf("%T \n", y)
	fmt.Printf("%b \n", y)
	fmt.Printf("%#x \t %b \t %x \n", y, y, y)

	// Sprint(), prints to a string
	s := fmt.Sprintf("%#x \t %b \t %x \n", y, y, y)
	fmt.Println(s)
}
