package main

import (
	"fmt"
	"runtime"
)

// go routines can comminicate between each other by means of channels
// Channels are typed, the data passed between routines must be of the
// same type

func main() {

	n := 10

	// Create 2 channels
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			done <- true
		}()

	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	//
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	for n := range c {
		fmt.Println(n)
	}

}
