package main

import (
	"fmt"
	"log"
)

func main() {
	msg, ok := hello("ajay", 28, 55)
	if !ok { // if ok == false
		log.Println("process failed", msg)
		return
	}
	fmt.Println(msg)
}

func hello(name string, age, marks int) (string, bool) {
	if name == "" {
		return "please provide a name", false // it will stop the current func and return values
	}
	if age == 0 {
		return "please provide your age", false
	}
	if marks == 0 {
		return "please provide your marks", false
	}
	fmt.Println(name, age, marks)
	return "success", true
}
