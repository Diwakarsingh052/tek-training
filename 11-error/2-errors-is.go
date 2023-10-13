package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("not in the root directory")

func main() {
	_, err := openFile("test.txt")

	if err != nil {
		//errors.Is looks inside the chain and check if custom err happened or not
		ok := errors.Is(err, ErrFileNotFound)
		if ok {
			fmt.Println("our custom error is found in the chain, lets create the file")
			//create the file
			//retry the operation
			//if still failed then stop this func
			return
		}
		log.Println(err)
		return
	}
	fmt.Println(err)
}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		//%w wrap the additional error context message
		return nil, fmt.Errorf("%w %w", err, ErrFileNotFound)
	}
	return f, nil
}
