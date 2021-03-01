package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	var counter int64
	const gs = 10

	var wg sync.WaitGroup
	wg.Add(gs)

	// Here we create a race condition, we have multipe go routines accessing
	// and modifying the same value

	// In this case we create we use the atomic pacage to prevent a race condition

	for i := 0; i < gs; i++ {
		go func() {
			// increment counter, by one using atomic
			atomic.AddInt64(&counter, 1)
			// time.Sleep(time.Second)
			// Gosched(), yields the processor alowing other goroutines to run.
			// it does not suspend the current goroutine, so execution resumes automatically
			// runtime.Gosched(), simply says go ahead and run something else
			runtime.Gosched()
			// Read counter value from atomic
			fmt.Println("counter:\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()
		fmt.Println("Coroutines:\t", runtime.NumGoroutine())
	}

	// Wait until all the goroutines are done, before exiting
	wg.Wait()
	// fmt.Println("Coroutines:\t", runtime.NumGoroutine())
	fmt.Println("final count:", counter)

	/** to see race conditions int programe run**/
	/**** go run -race main.go ***/

}
