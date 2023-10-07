package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	//b := make([]int, 3, 20)  // only three elem max can be copied here as that is the available len only
	b := make([]int, len(x[1:4]), 20) // creating a new backing array
	copy(b, x[1:4])                   //
	b[0] = 999
	fmt.Println(b)
	fmt.Println(x)
}
