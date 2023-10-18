package main

import (
	"fmt"
)

// model data using structs
type user struct {
	Name     string `json:"name_something"`
	Password string `json:"-"`
	age      int
	marks    []int
}

// func (receiver) methodName(args)returnType {}
func (u *user) show() {
	fmt.Println(u)
}
func (u *user) updateAge(age int) {
	u.age = age
}

func main() {
	//l := log.New(os.Stdout, "sales:", log.Lshortfile)
	//l.Println()
	u1 := user{
		Name:     "raj",
		Password: "abc",
		age:      30,
		marks:    nil,
	}

	u1.show()
	u1.updateAge(31)
	u1.show()
}
