package main

import (
	"fmt"
	"sync"
	"time"
)

var cab int = 1

func main() {
	var wg = &sync.WaitGroup{}
	//var m = &sync.Mutex{}
	names := []string{"a", "b", "c", "d"}
	for _, name := range names {
		wg.Add(1)
		go bookCab(name, wg)

	}
	wg.Wait()
}

func bookCab(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)
	//m.Lock()
	//critical section where we are using a shared resource
	// when a goroutine acquires a lock then another go routine can't access the critical section
	//until the lock is not released
	//any read , write from other goroutines would not be allowed after lock is acquired
	//defer m.Unlock()
	if cab >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		cab--
	} else {
		fmt.Println("car is not available for", name)
	}
	fmt.Println()
}
