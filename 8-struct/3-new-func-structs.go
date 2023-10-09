package main

import (
	"fmt"
	"log"
)

type Conf struct {
	db string // unexported field
}

//change the signature to return a pointer to conf
// if no dataSourceName is provided then log a message and return otherwise assign dataSourceName to db

func NewConf(dataSourceName string) *Conf {

	if dataSourceName == "" {
		log.Println("please provide datasource name")
		return nil
	}

	return &Conf{db: dataSourceName}
}

func main() {
	c := NewConf("postgres")
	//sql.Open()
	//log.New() // design pattern
	//os.OpenFile()
	fmt.Println(c)
}
