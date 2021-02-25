package main

import "fmt"

// A struct is a composite data type of aggregate data type or complex data type
// It allows us to compose together values of different types
type person struct {
	firstname string
	lastname  string
	age       int
}

type secretAgent struct {
	// Embedded struct
	// You do not give and embeded struct a type name
	person
	licenseToKill bool
	age           string // Name collision
}

func main() {
	// P1 is a value of type person
	// person{} is a composite literal
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

	sa1 := secretAgent{
		person: person{
			firstname: "Jason",
			lastname:  "Bourne",
			age:       32,
		},
		licenseToKill: true,
		age:           "thirty-two",
	}

	fmt.Println(sa1)
	fmt.Println("age name causes a colision")
	fmt.Println(sa1.person.age)
	fmt.Println(sa1.age)

}
