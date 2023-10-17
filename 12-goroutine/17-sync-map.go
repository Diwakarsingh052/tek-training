package main

import (
	"fmt"
	"sync"
)

func main() {
	// Initialize a new map
	var m sync.Map

	// Store: Set key/value pairs
	m.Store("animal", "giraffe")
	m.Store("color", "orange")

	// Load: Get a value
	if v, ok := m.Load("animal"); ok {
		fmt.Println("Loaded value:", v) // Prints: "Loaded value: giraffe"
	}

	// LoadOrStore: Returns existing value for a key or store a new value
	if v, loaded := m.LoadOrStore("color", "green"); loaded {
		fmt.Println("Found existing value:", v) // Prints: "Found existing value: orange"
	} else {
		fmt.Println("Stored new value:", v) // This won't run in this case
	}

	// Delete: Remove a key/value pair
	//m.Delete("animal")

	// Use Load to show that "animal" was deleted
	if _, ok := m.Load("animal"); !ok {
		fmt.Println("Value not found") // Prints: "Value not found"
	}

	// Range: Iterating over the map
	m.Range(func(k, v interface{}) bool {
		fmt.Println("Key:", k, "Value:", v)
		return true
	})

	//The Range function of sync.Map uses a function as its parameter. The function passed to Range is expected to take two parameters, a key and a value, and return a boolean.
	//The boolean returned by this function controls whether the iteration process should continue or not.
	//The part return true in your function means you want to continue to the next item in the sync.Map. If the function returned false, then the Range function would stop iterating through the sync.Map.
	//So, in the context of your code, returning true means "I'm done processing the current key-value pair, continue to the next pair". If you returned false, it would mean "I'm done processing the current key-value pair, and I don't need to see any more pairs".
	//This design gives you the power to stop the iteration early if you've found what you're looking for, or if you need to abort the process for some reason.

}
