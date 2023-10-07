package main

import "fmt"

func main() {

	var i []int // nil

	//preallocate the backing array
	i = make([]int, 0, 20) //make(type,len,cap)
	i = append(i, 10, 20, 30)
	inspectSlice("i", i)

}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
