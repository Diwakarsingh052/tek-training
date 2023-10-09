package main

import (
	"fmt"
	"log"
	"strconv"
)

func SumString(s, x string) (int, error) { // err must be the last value to be returned

	a, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	b, err := strconv.Atoi(x)
	if err != nil {
		return 0, err
	}

	return a + b, nil //success

}

func main() {
	sum, err := SumString("10", "0")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(sum)

}
