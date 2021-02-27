package main

import "fmt"

//
func main() {
	// Rankings map
	rm := map[string]int{
		"england":  11,
		"zimbabwe": 50,
		"zambia":   67,
	}

	fmt.Println(rm)
	fmt.Println(rm["england"])

	// comma ok idiom, lets you check a field exists before accessing
	v, ok := rm["zaire"]
	fmt.Println(v)
	fmt.Println(ok)

	// we can check a value exists using an if statement
	if v, ok := rm["zaire"]; ok {
		fmt.Println(v)
	}

	if v, ok := rm["england"]; ok {
		fmt.Println(v)
	}
}
