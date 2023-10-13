package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{} // We create a wait group.
	// A wait group waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of goroutines to wait for.
	//Then it calls Wait to block
	// until all goroutines have finished.

	for i := 1; i <= 10; i++ {
		wg.Add(1) // Here we increment the WaitGroup counter by one.
		// It means we have one more go routine into the counter.

		go work(i, wg) // Create a new goroutine to perform some work.
	}
	wg.Wait() // We block here until all the goroutines call Done()
	// and respectively decrement the counter so it reaches zero.
}

func work(i int, wg *sync.WaitGroup) {
	defer wg.Done() // This will be called when the function exits.
	// It decrements the WaitGroup counter by one, signifying this routine's work is done.

	wg.Add(1) // Here we increment the WaitGroup counter by one again.
	// It means we have one more nested process to wait for.

	go func(id int) { // Create a new inner goroutine.
		defer wg.Done() // This will be called when the inner routine exits.
		// It decrements the WaitGroup counter by one, signifying this routine's work is done.
		fmt.Println("work", id)
	}(i)
}
ʼ