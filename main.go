package main

import "fmt"

func main() {
	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)

}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprintln(fn, ln+`, says "hello"`)
	b := false
	return a, b
}
