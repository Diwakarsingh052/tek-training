package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wgWorker := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {

			wgWorker.Add(1)
			go func(i int) {

				defer wgWorker.Done()
				ch <- i

			}(i)
		}

		wgWorker.Wait()
		close(ch)

	}()

	go func() {

		defer wg.Done()

		for r := range ch {
			fmt.Println("recieved", r)
		}

	}()

	wg.Wait()

}
