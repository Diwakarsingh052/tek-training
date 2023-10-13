package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	// To keep track of if the goroutine work is finished or not,
	// we need another goroutine to close the channel (done).
	// This pattern is useful in case of multiple producers sending data on
	// the channel and single consumer reading from it. Closing the channel
	// signals the consumer about the completion of data sending by all producers.
	wgWorker := sync.WaitGroup{}

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		time.Sleep(4 * time.Second)
		c1 <- "1"
		c1 <- "4"
	}()
	go func() {
		defer wgWorker.Done()
		time.Sleep(2 * time.Second)
		c2 <- "2"
	}()
	go func() {
		defer wgWorker.Done()
		c3 <- "3"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done) // We are closing the channel (done) when all goroutines are finished sending.
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Infinite for loop with select is useful here because we don't know how many value
		// we would receive from the channels
		// whatever channel sends the data first, that case will execute.
		// This gives us concurrency and we don't have to wait for a particular channel to finish.
		// Also, it allows the easy addition of more channels in the future.
		for {
			// whichever case is not blocking exec that first
			//whichever case is ready first exec that.
			// possible cases are chan recv , send , default
			select {
			case x := <-c1:
				fmt.Println(x)
			case x := <-c2:
				fmt.Println(x)
			case x := <-c3:
				fmt.Println(x)
			case <-done: // this case will execute when channel is closed, signalling all work done
				fmt.Println("work is finished")
				return
			}
		}
	}()
	wg.Wait()
}
