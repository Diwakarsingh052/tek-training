package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("end of the main")
}

func greet() {
	//2nd
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("please provide name , age , marks")
		//os.Exit(1)
		return // stops the exec of the current func
	}

	name := data[0]
	ageString := data[1]
	marksString := data[2]
	//var err error // default value is nil // indicates no error
	//fmt.Println(err)
	//handle the error in the next line
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println("some kind of error", err)
		return

	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(name, age, marks)
	fmt.Println()

}
