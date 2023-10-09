package main

import "fmt"

// model data using structs
type user struct {
	Name     string // fields
	Password string
	age      int
	marks    []int
}

// model a book type // name, date , code, authors

func main() {

	var u user // setting all fields to the default values // user{}

	u.Name = "ajay"
	u1 := user{
		Name:     "raj",
		Password: "raj",
		age:      30,
		marks:    []int{10, 20, 40, 50},
	}

	fmt.Printf("%#v\n", u1) // extra info about the types
	fmt.Printf("%+v\n", u1) // field:value
	fmt.Println(u)
}
