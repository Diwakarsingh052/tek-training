package main

import (
	"log"
	"os"
)

func main() {
	// Attempt to open the file "text.txt", using `os.Open()`. This function returns two values: a pointer to a `File` object, and an `error`.
	f, err := os.Open("text.txt")

	// `defer` is used to ensure that `f.Close()`, which is called to close the file, is executed after all other code in the function has finished running.
	// Regardless of what happens inside the function, the file will be closed before the function ends.
	// This is especially useful to prevent resource leaks in case of an error occurring anywhere in the function because `defer` runs whether the function exits naturally or through a panic.
	defer f.Close()

	// If opening the file resulted in an error, it is logged, and the program returns, which stops further execution.
	// Thanks to `defer`, we know that the file will be closed in any event, even if an error occurs.
	if err != nil {
		log.Println(err)
		return
	}

	// Further processing of the file would occur here...

}
