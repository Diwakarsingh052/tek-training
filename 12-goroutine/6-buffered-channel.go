package main

import (
	"fmt"
	"sync"
)

// A send on a buffered channel can proceed if there is room in the buffer.

func main() {
	// Define a wait group
	var wg sync.WaitGroup

	// Define a buffered channel that can hold 2 integers
	ch := make(chan int, 2)

	// Increment wait group counter
	wg.Add(2)

	// Start a goroutine to send data into the channel
	go func() {
		// Decrement wait group counter upon completion
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Printf("Sending: %d\n", i)
			// Write to the channel
			ch <- i // sender only needs space in buffer, if space is available to put the value,
			// then it would proceed // it doesn't care about receiver.
		}
	}()

	// Start another goroutine to read data from the channel
	go func() {
		// Decrement wait group counter upon completion
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			// Read from the channel
			val := <-ch
			fmt.Printf("Received: %d\n", val)
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()

	// Close the channel
	close(ch)

	// Attempt to read from the closed channel
	val, ok := <-ch
	fmt.Printf("Read from closed channel, val: %d ok: %t\n", val, ok)
}
