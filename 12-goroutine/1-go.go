package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// counter to add number of goroutines
	// that we are starting or spinning up
	wg.Add(1)
	go hello(wg) // spawning a goroutine
	fmt.Println("end of the main")

	wg.Wait() //block until the counter is reset to 0
	//bad way of doing it
	//time.Sleep(2 * time.Second)
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter
	fmt.Println("hello from the hello func")
}
