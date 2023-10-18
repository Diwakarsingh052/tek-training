package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define a slice of strings, i.e., names
	names := []string{"sangmish", "somshekhar", "sidarth", "harshita", "nikita", "bhumika", "purvi",
		"aftab", "jeevan", "vishal", "dhanush", "santosh", "sharad", "shashi", "sindhu", "mahesh",
		"keerthi", "harshath", "ashwini",
		"vaurn", "sahil", "sreedhar", "surya", "tejaswani", "vikalp", "pragalb", "satyam", "sandeep"}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Pick a random name from the slice
	randomName := names[rand.Intn(len(names))]

	fmt.Println("Randomly picked name:", randomName)
}
