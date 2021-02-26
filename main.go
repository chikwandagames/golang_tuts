package main

import (
	"fmt"
)

// INTERFACES
// Interfaces allow us to define behavior and POLYMORPHISM

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	licenceToKill bool
}

type human interface {
	// Any type that implements laugh() is also now of type human
	// Therefor a value can be of more than one type
	laugh()
}

// POLYMORPHISM, the laugh method behaves differently depending on
// type attached to it
func (p person) laugh() {
	fmt.Printf("%T LOL \n", p)
}

func (s secretAgent) laugh() {
	fmt.Printf("%T LMAO \n", s)
}

func isAllowedToShoot(h human) {
	// Here we switch by type
	// Assertion h.(person).first, assert that h is of type person,
	// Assertion gets you back to the concrete type, from human to person in this case
	switch h.(type) {
	case person:
		fmt.Printf("%v is NOT allowed to shoot\n", h.(person).first)
	case secretAgent:
		fmt.Printf("%v is Allowed to shoot\n", h.(secretAgent).first)
	default:
		fmt.Println("No No No")
	}

}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		licenceToKill: true,
	}

	p1 := person{
		last:  "Bourne",
		first: "Jason",
	}

	p1.laugh()
	sa1.laugh()
	isAllowedToShoot(p1)
	isAllowedToShoot(sa1)

}
