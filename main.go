package main

import "fmt"

func main() {
	// DONT use fall through
	switch {
	case false:
		fmt.Println("Nope")
	case (2 == 4):
		fmt.Sprintln("Nope")
	case (3 == 5):
		fmt.Println("nope")
		fallthrough
	case (3 == 3):
		fmt.Println("yes")
		fallthrough
	case (4 == 4):
		fmt.Println("yes")
	default:
		fmt.Println("default")
	}

	n := "Bond"
	switch n {
	case "Wisdom":
		fmt.Println("nope")
	case "Bond", "trev": // multipe cases
		fmt.Println("james")
	default:
		fmt.Println("default")
	}

}
