package main

import "fmt"

func main() {
	var p *int
	a := 10
	p = &a

	fmt.Println(p)
	fmt.Println(&a)

	*p = 20 // dereferencing operator // value at this address
	fmt.Println(a)
}
