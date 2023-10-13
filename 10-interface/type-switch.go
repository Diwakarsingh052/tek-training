package main

import "fmt"

type student struct {
}

func main() {
	var s student
	checkType(10)
	checkType(s)
	checkType("hello")
	checkType(true)
}

func checkType(i any) {
	switch i.(type) {
	case int:
		fmt.Println("it is int")
	case string:
		fmt.Println("it is string")
	case student:
		fmt.Println("it is a struct student")
	default:
		fmt.Println("nothing matches")

	}
}
