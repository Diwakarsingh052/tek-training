package main

import "fmt"

type user struct {
	name string
	age  int
}

type movies struct {
	movieName string
	// anonymous field which has no name
	// anonymous field gets there names from the types
	user // embedding
}

func main() {
	m := movies{
		movieName: "abc",
		user: user{
			name: "xyz",
			age:  29,
		},
	}

	fmt.Println(m.age)
}
