package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type admin struct {
	user // embedding
	role []string
}
type book struct {
	user     // embedding
	bookName string
	pages    int
}
type student struct {
	u     user // not embedding // this is a field of user type
	class int
	marks int
}

func main() {

	b := book{
		user: user{
			name:  "raj",
			email: "raj@email.com",
		},
		bookName: "learn Go",
		pages:    400,
	}

	s := student{
		u: user{
			name:  "ajay",
			email: "ajay@email.com",
		},
	}

	s.u.updateEmail("aja@student.com")
	fmt.Println(s)
	b.user.updateEmail("raj@gmail.com")
	fmt.Println(b)

}
