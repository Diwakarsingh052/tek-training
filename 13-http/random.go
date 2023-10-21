package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define a slice of strings, i.e., names
	names := []string{"sangmish", "somshekhar", "sidarth", "harshita", "nikita", "bhumika", "purvi",
		"aftab", "jeevan", "vishnu", "vishal", "dhanush", "santosh", "sharad", "shashi", "sindhu", "mahesh",
		"keerthi", "harshath", "ashwini",
		"vaurn", "sahil", "sreedhar", "surya", "tejaswani", "vikalp", "pragalb", "satyam", "sandeep"}

	rand.Seed(time.Now().UTC().UnixNano())

	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Pick a random name from the slice
	randomName := names[rand.Intn(len(names))]

	fmt.Println("Randomly picked name:", randomName)
}

func shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

	fmt.Println(slice)
}
