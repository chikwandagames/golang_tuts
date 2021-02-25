package main

import "fmt"

// we define functions with parameters
// we call our function and pass in arguments
// everythin in go is PASS BY VALUE

func main() {
	printName("wisdom")

	s1 := foo("David")
	fmt.Println(s1)
}

// func (r receiver) ideintifire(parameters) (return(s)) { ... }
func printName(s string) {
	fmt.Printf("Hello %v\n", s)
}

func foo(st string) string {
	// StringPrint, prints to a string
	return fmt.Sprint("Hello from ", st)
}
