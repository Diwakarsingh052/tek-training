package main

import (
	"fmt"
	"tek-training/db"
)

func main() {
	// call the NewConf
	//try to change the db connection
	db.Open("postgres")
	fmt.Println(db.Conn)
	db.Conn = "mysql"
	fetchData(db.Conn)
}

func fetchData(conn string) {
	fmt.Println("fetching data from", conn)
}
