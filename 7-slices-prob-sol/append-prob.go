package main

import "fmt"

//If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array
//that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array.

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := x[3:6] // index:len  Length 3 Cap 4

	inspectSlice("b", b)
	b = append(b, 777, 888, 999)
	inspectSlice("b", b)
	inspectSlice("x", x)

}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
