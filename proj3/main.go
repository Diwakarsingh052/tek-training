package main

import (
	"proj3/stores"
	"proj3/stores/mysql"
	"proj3/stores/postgres"
)

func main() {
	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}
	m := mysql.New("mysql")
	p := postgres.New("postgres")
	//assigning mysql instance to storer
	// we can do this because mysql impls all the methods
	// of the storer interface
	//stores.StorerInterface = m
	//stores.StorerInterface.Create(u)

	ms := stores.NewService(m)
	ms.Create(u)

	ps := stores.NewService(p)
	ps.Create(u)
}
