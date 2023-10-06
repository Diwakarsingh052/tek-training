package main

import (
	"fmt"
	"tek-training/sum"
)

func main() {
	fmt.Println()
	sum.Add(2, 3)
	sum.Add(2, 5)

	//fmt.Sprintf()  // look for design patterns for unexported functions
	// create a package named as greet // create an exported func name as hello
	// hello(name) -> hello name
}
