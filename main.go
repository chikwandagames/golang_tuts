package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs \t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	counter := 0
	const gs = 10

	var wg sync.WaitGroup
	wg.Add(gs)

	// Create a mutex
	var mu sync.Mutex

	// Here we create a race condition, we have multipe go routines accessing
	// and modifying the same value

	// MUTEX
	// To prevent a race condition, we need to lock the shared variable down, so that while
	// on go routine is accessing the variable no other go routine can access that value

	for i := 0; i < gs; i++ {
		go func() {
			// Here we use a mutex to lock, so from the line of code below, while a goroutine
			// is accessing a shared variable not other goroutine can access that variable until its
			// unlocked
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			// Gosched(), yields the processor alowing other goroutines to run.
			// it does not suspend the current goroutine, so execution resumes automatically
			// runtime.Gosched(), simply says go ahead and run something else
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Coroutines:\t", runtime.NumGoroutine())
		fmt.Println("count:", counter)
	}

	// Wait until all the goroutines are done, before exiting
	wg.Wait()
	// fmt.Println("Coroutines:\t", runtime.NumGoroutine())
	fmt.Println("final count:", counter)

	/** to see race conditions int programe run**/
	/**** go run -race main.go ***/

}
