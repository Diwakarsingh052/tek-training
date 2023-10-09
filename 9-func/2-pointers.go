package main

import "fmt"

func main() {
	i := 10
	p := &i
	//p := &i // var p *int
	//fmt.Println(p)
	//fmt.Println(&i)
	update(p)
	fmt.Println(i)
	fmt.Println(&p, "p")
}
func update(ptr *int) {
	*ptr = 10000
	fmt.Println(&ptr, "ptr")
}
