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

}
