package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
//send will block until there is no recv

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) //unbuffered channel
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch <- 20 // send will block until no receiver is ready
	}()
	go func() {
		defer wg.Done()
		x := <-ch //it is a blocking call until we don't recv the value
		fmt.Println(x)
	}()

	wg.Wait()
	fmt.Println("end of main")

}
