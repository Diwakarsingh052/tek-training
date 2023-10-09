package main

import "fmt"

type user struct {
	name string
	age  int
}

type movies struct {
	movieName string
	// anonymous field which have no name
	// anonymous field get there names from the types
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
