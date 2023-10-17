package main

import (
	"sync"
)

var myMap = make(map[int]int)
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			myMap[n] = n * n

		}(i)
	}
	wg.Wait()
}
