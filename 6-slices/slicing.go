package main

import "fmt"

func main() {
	//slicing // reference
	a := []int{10, 20, 30, 40, 50}
	b := a[1:4] // index:len
	b = a[:]    // take the whole slice
	b = a[:3]   // start from 0 index till the length provided
	b = a[2:]   // from the 2nd index till the end
	fmt.Println(b)
}
