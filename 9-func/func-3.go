package main

import "fmt"

// passing function to another function
func main() {
	//this function would not be called at below line
	//we are defining the working of the func
	//add := func(x, y int) {
	//	fmt.Println(x + y)
	//}
	sub := func(x, y int) int {
		//fmt.Println(x - y)
		return x - y

	}

	//add could be passed to operate function because the signature matches
	operate(add, 100, 20)
	operate(sub, 100, 20)
	//add two more functions for multiply and division

}

func add(x, y int) int {
	//fmt.Println(x + y)
	return x + y

}

// change the signature to accept the float values
func operate(op func(a int, b int) int, s, y int) {
	fmt.Println(op(s, y))
}
