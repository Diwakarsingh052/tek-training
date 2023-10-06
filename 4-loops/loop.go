package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 0; i <= 10; i++ {

	}

	i := 0
	for ; i <= 10; i++ {

	}

	i = 0
	for i <= 10 {
		//do work here

		i++

	}

	//ranging over a list of strings
	for index, value := range os.Args {
		fmt.Println(index, value)
	}
}
