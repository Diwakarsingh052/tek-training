package postgres

import (
	"fmt"
	"proj3/stores"
)

type Conn struct {
	db string
}

func New(db string) Conn {
	return Conn{db: db}
}

//implementing the Storer interface over the conn struct

func (p Conn) Create(usr stores.User) error {
	fmt.Println("adding to postgres", usr)
	return nil
}
func (p Conn) Update(usr stores.User) error {
	fmt.Println("updating in postgres", usr)
	return nil
}

func (p Conn) Delete(usr stores.User) error {
	fmt.Println("deleting in postgres", usr)
	return nil
}
