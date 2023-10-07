package main

import "fmt"

func main() {
	// slice points to an array (backing , underlying) in the memory

	//var i []int // nil
	i := []int{10, 20, 30, 40}
	// this will cause panic as length is not available to store the value
	i[6] = 100 //update ops // it would not grow the slice
	fmt.Println(i)
	if i == nil {
		fmt.Println("it is nil slice")
	}

	//for i,v := range slice {}
}
