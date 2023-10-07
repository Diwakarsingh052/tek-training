package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)
	//slicing
	cmdArgs := os.Args[1:]
	fmt.Println(cmdArgs[0])

}

// wap that takes cmd args and check whether 4 values are provided or not
// progName, name,age,marks
//if values are not there quit the program // os.Exit(1)
