package main

import "fmt"

func main() {

	store := make(map[int]*int)
	a := []int{10, 20, 30, 40}
	for i, v := range a {
		v := v
		store[i] = &v
	}
	for k, v := range store {

		fmt.Println(k, *v)
	}
}
