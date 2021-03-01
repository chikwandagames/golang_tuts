package main

import (
	"fmt"
	"time"
)

// go routines can comminicate between each other by means of channels
// Channels are typed, the data passed between routines must be of the
// same type

func main() {

	// Channel declaration, unbufferd
	c := make(chan int)
	// 	c2 := make(chan int, 10), a channel that can take 10 ints

	// Here we have 2 goroutines accessing the same value, so they need to wait
	// for each other, on goroutine can only access the value c when the othe has
	// finished using the value

	go func() {
		for i := 0; i < 10; i++ {
			// Feed i into a the channel
			// At this point code execution stops until something takes the value
			// out of the channel i.e. <-c
			c <- i
		}
	}()

	go func() {
		for {
			// fmt.Println(<-c), this receives data from a channel, and prints
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)

}
