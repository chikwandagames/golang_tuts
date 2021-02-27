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

	//
	rm["todd"] = 33

	for k, v := range rm {
		fmt.Printf("%v: %v\n", k, v)
	}

}
