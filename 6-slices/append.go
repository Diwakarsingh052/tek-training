package main

import "fmt"

// https://go.dev/ref/spec#Appending_and_copying_slices
/*
	append func working

	i := []int{10, 20, 30, 40, 50 } // len = 5 , cap =5
	append(i,60) // not enough cap so allocation is going to happen

//  sufficiently large underlying array.
	underlying array -> [10 20 30 40 50,60,{},{}] len =6 cap = 8

append(i,70,90,300) // allocation would happen as we don't have enough cap to fit three values
	underlying array -> [10 20 30 40 50,60,70,80,90, , , , ] len =9 cap = 13

	If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values.
    Otherwise, append re-uses the underlying array.
*/
func main() {

	var a []int
	b := []int{89, 69, 798}
	a = append(a, b...) // ellipsis // unpack your slice
	fmt.Println(a)

	//cap // how many elems backing array can contain
	// len // how many elems slice is referring to in backing array

	i := []int{10, 20, 30}
	fmt.Println("cap", cap(i), len(i))
	i = append(i, 40)
	fmt.Println("cap", cap(i), len(i))
	i = append(i, 50, 60, 70, 80)
	fmt.Println("cap", cap(i), len(i))

}
