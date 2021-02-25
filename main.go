package main

import "fmt"

func main() {
	// Anonymous struct
	p1 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "Jame",
		lastname:  "Bond",
		age:       32,
	}

	fmt.Println(p1)
	fmt.Println(p1.firstname)

}
