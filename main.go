package main

import (
	"fmt"
)

// go routines can comminicate between each other by means of channels
// Channels are typed, the data passed between routines must be of the
// same type

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		// feed true into the done channel
		done <- true

	}()

	go func() {

		for s := 0; s < 10; s++ {
			c <- s
		}
		done <- true

	}()

	go func() {
		// 2 channels waiting for a value to come through
		// <- means Read data from the channel done
		// this basically throws a value away
		// x := <- done, world save the value in x
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
