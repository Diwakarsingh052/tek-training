package main

import "fmt"

func main() {
	a := 10 // x80
	var p *int = &a
	update(p)
	fmt.Println(a)
}

func update(p *int) { // x80 would be copied to p // p var would have its own address // p address = x120
	*p = 20
}
