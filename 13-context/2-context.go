package main

import (
	"context"
	"fmt"

	"sync"

	"time"
)

func main() {

	// value returned is a pointer to a newly allocated zero value of that type
	wg := new(sync.WaitGroup)

	// gives an empty container to put context values
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)

	//it would cleanup the resources taken up context
	defer cancel() // it cancels the context immediately

	ch := make(chan int)

	wg.Add(1)
	go func() {

		defer wg.Done()
		//sending the value to the channel if timout is not over
		select {
		case ch <- 20:
			fmt.Println("values sent")
			//ctx.Done evaluates when timeout happens
		case <-ctx.Done():
			fmt.Println("receiver not present")
		}

	}()
	recv(ctx, ch)
	wg.Wait()

}

func recv(ctx context.Context, ch chan int) {
	time.Sleep(time.Second)
	select {
	//trying to receive value if it is available
	case val := <-ch:
		fmt.Println(val)

		//if ctx.done evaluates, which means another goroutine already
		//exited and no longer to receive the values
	case <-ctx.Done():
		fmt.Println(":", ctx.Err())

	}

}
