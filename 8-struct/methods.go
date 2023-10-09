package main

import "fmt"

//https://go.dev/doc/faq#methods_on_values_or_pointers

// model data using structs
type user struct {
	Name     string // fields
	Password string
	age      int
	marks    []int
}

func (u *user) show() { // func(receiver)methodName(Args)returnTypes {}
	fmt.Printf("%+v\n", u)
}
func (u *user) updatePassword(password string) {
	u.Password = password
}

func main() {
	u1 := user{
		Name:     "raj",
		Password: "raj",
		age:      30,
		marks:    []int{10, 20, 40, 50},
	}
	u1.updatePassword("abc")
	u1.show()
}
