package main

import (
	"fmt"
	"time"
)

// GORUTINES
// Goroutines are like very lightweight threads managed by the go runtime
// they enable us to create asychronous parrallel programs

// Go routines is a way to turn a sequential program into a concurrent program
// without the need for threads

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	// Addging the go key word changes these calls to goroutines
	// Now these calls will excecute Asynchronously
	go compute(5)
	go compute(5)

	// this is just a hack to prevent the main function exiting
	// before the goroutines have completed execution
	fmt.Scanln()
}
