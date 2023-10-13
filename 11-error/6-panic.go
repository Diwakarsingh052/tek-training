package main

import "fmt"

// Define the main function, which is the entry point of our application.
func main() {
	// Define a list of deferred functions. "defer" means that the function's execution is postponed until the surrounding function (i.e., "main") returns.
	// Defer maintains a stack, meaning functions are executed in Last-In-First-Out (LIFO) order. Therefore, 'fmt.Println(1)' will execute last.
	defer fmt.Println(1)

	// 'fmt.Println(2)' will be the second to the last function to execute.
	defer fmt.Println(2)

	// 'fmt.Println(3)' will be the first deferred function to execute because of the LIFO order.
	defer fmt.Println(3)

	// The 'panic' function throws an unhandled runtime error, which stops execution of the current function (i.e., "main").
	// After a panic, execution stops and all deferred functions in the stack are executed. Once all deferred functions are done, the program stops running.
	panic("some kind of msg")

	// Declare an integer slice named "i". However, its length is 0.
	var i []int

	// Trying to access the 101st element and set its value to 1000 would cause an "index out of range" error. However, this line is never reached because of the panic.
	i[100] = 1000
}
