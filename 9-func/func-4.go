package main

import "fmt"

// creating own custom type
type money int

type ops func(a int, b int)

func main() {
	//var dollar money = 1000 // underlying type of money is int
	//time.Duration // custom type defined by the standard lib
	sub := func(x, y int) {
		fmt.Println(x - y)
	}
	operate(sub, 100, 90)
}

func operate(op ops, s, y int) {
	op(s, y)
}
