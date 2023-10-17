package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = &sync.WaitGroup{}
	var wgWorker = &sync.WaitGroup{}
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	done := make(chan struct{})

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		time.Sleep(time.Second * 3)
		c1 <- "one" // send
	}()
	go func() {
		defer wgWorker.Done()
		time.Sleep(time.Second)
		c2 <- "two" // send
	}()

	go func() {
		defer wgWorker.Done()
		c3 <- "three" // send
		c3 <- "four"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done)

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case x := <-c1:
				fmt.Println(x)
			case y := <-c2:
				fmt.Println(y)
			case z := <-c3:
				fmt.Println(z)
			case <-done:
				fmt.Println("all the values are recvd")
				return

			}
		}

	}()

	wg.Wait()
}
