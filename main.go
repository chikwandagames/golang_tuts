package main

import "fmt"

func main() {

	// Because foo() is defered, it will only be executed after main() exits
	// We can use this to close, a file after a program finishes execution
	defer foo()
	bar()

}

func foo() {

	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
