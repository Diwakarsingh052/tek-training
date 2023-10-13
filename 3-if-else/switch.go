package main

import "fmt"

func main() {

	ext := ".txt"
	// break is implicit
	// write fallthrough when needed a behaviour of going to the next case
	switch ext {
	case ".txt":
		fmt.Println("a text file")
		fallthrough
	case ".jpeg":
		fmt.Println("an image")

	default:
		fmt.Println("unknown file type")

	}
}
