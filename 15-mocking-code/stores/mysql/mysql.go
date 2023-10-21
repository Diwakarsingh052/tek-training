package mysql

import (
	"database/sql"
	"fmt"
	"mocking-cod/stores"
)

type Conn struct {
	db *sql.DB
}

// NewConn setups the db connection
func NewConn(db *sql.DB) *Conn {
	return &Conn{db: db}
}

func (p *Conn) Create(usr core.User) error {
	fmt.Println("adding to mysql", usr)
	return nil
}
func (p *Conn) Update(usr core.User) error {
	fmt.Println("updating in mysql", usr)
	return nil
}
func (p *Conn) Delete(usr core.User) error {
	fmt.Println("deleting in mysql", usr)
	return nil
}
