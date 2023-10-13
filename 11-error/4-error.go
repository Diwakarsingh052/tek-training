package main

import (
	"errors"
	"fmt"
)

//Error should be suffixed in the name of the struct use for error handling

type QueryError struct {
	Func  string
	Input string
	Err   error
}

// implementing the error interface // format your msg here
func (q *QueryError) Error() string {
	return "main." + q.Func + ": " + "input " + q.Input + " " + q.Err.Error()
}

var ErrNotFound = errors.New("not found")
var ErrMismatch = errors.New("mismatch")

func SearchSomething(s string) error {

	//this is the case when data is not present
	return &QueryError{
		Func:  "SearchSomething",
		Input: s,
		Err:   ErrNotFound,
	}
}

func main() {
	var q *QueryError // nil
	//os.PathError{}
	//os.LinkError{}
	err := SearchSomething("whatever")
	fmt.Println(err)
	err = SearchSomething("abc")
	fmt.Println(err)
	if err != nil {
		if errors.As(err, &q) {
			fmt.Println("custom error found in the chain", q.Func)
			return
		}
		return
	}
}
