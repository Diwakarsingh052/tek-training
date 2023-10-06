package main

import "fmt"

func main() {
	//n, err := fmt.Println()
	if n, err := fmt.Println(); n == 1 {
		fmt.Println("nothing was written")
	} else {
		fmt.Println(n, err)
	}

	var n string = "hello"
	fmt.Println(n)

	//fmt.Println(n) // n scope ended in if block

}
