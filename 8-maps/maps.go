package main

import "fmt"

func main() {
	//map[key]value
	dictionary := make(map[string]string)

	dictionary["up"] = "above"
	dictionary["below"] = "down"

	fmt.Println(dictionary["up"])
	for _, value := range dictionary {
		fmt.Println(value)
	}
	//Passing a map to the len function tells you the number
	//of key-value pairs in a map.
	fmt.Println(len(dictionary))
}
