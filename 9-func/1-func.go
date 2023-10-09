package main

import "fmt"

type operation func(x, y int) int

func main() {
	s := func(a, b int) int {
		return a - b
	}
	operate(add, 20, 30)
	operate(s, 40, 30)

}
func operate(op operation, a, b int) {
	fmt.Println(op(10, 20))
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

//create a func for sub it takes two int and do the subtraction,
//and it returns output in int
