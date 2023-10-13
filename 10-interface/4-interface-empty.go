package main

import "fmt"

type user struct {
	name string
}

func main() {
	var i interface{} = "str"
	//var a any // any is an alias to interface{}
	i = 10
	i = true
	i = "hello"
	var u user
	i = u
	str, ok := i.(user) // type assertion // checking if interface is storing user
	if !ok {
		fmt.Println("user type is not there")
		return
	}
	fmt.Println(str)
}
