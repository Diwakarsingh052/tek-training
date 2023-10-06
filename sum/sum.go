package sum

import "fmt"

func Add(a, b int) { // Add is an exported func, because the first letter is uppercase
	fmt.Println("calling add", a+b)
	doSomething()
}
