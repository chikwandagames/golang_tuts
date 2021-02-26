package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	licenceToKill bool
}

// METHOD
// (s secretAgent) is a Receiver
func (s secretAgent) speak() {
	fmt.Printf("I am %v %v \n", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		licenceToKill: true,
	}

	fmt.Println(sa1)

	sa1.speak()

}
