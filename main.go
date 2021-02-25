package main

import "fmt"

// A struct is a composite data type of aggregate data type or complex data type
// It allows us to compose together values of different types
type person struct {
	firstname string
	lastname  string
}

func main() {
	// P1 is a value of type person
	p1 := person{
		firstname: "Jame",
		lastname:  "Bond",
	}

	p2 := person{
		firstname: "Miss",
		lastname:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.firstname)

}
