package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	//wg.Add(10)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("work", id)
		}(i) //We pass `i`
		// to the anonymous function
		// in order for it to have a distinct copy of the variable
		// which is not affected by the concurrent execution of the loop.
	}
	wg.Wait()
}
