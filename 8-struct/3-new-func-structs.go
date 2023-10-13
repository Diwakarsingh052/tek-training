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

// NewConf function is created to create instance of the Conf struct
// any external package can use NewConf to have a db connection
// not db string is unexported, which means no one from the outside package could change it
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
