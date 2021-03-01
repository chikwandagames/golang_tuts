package main

import (
	"fmt"
	"sync"
)

// go routines can comminicate between each other by means of channels
// Channels are typed, the data passed between routines must be of the
// same type

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	// Add 2 for each wait group
	wg.Add(2)

	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {

		for s := 0; s < 10; s++ {
			c <- s
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
