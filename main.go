package main

import (
	"fmt"
)

func main() {
	// Create channel
	passwordChannel := make(chan string)
	// Listener channel
	listenerChannel :=make(chan string, 1)
	// // Slice of Words
	passwordSlice := Parser()
	// fmt.Println(passwordSlice)
	// fmt.Println(len(passwordSlice))

	rawPassword := "taylormade"
	hashedPasswordTest, _ := HashPassword(rawPassword)	
	// Create X amount of workers
	for i := 0; i < 100; i++ {
		go Worker(passwordChannel, hashedPasswordTest, listenerChannel )
	}

	// Sends a password to the channel
	for _, password := range passwordSlice {
		passwordChannel <- password
	}
	fmt.Println(<-listenerChannel)
}



// TO DO

// Parse top password txt into an array
// learn to export it.

// Create worker function
// worker function will take a password and compare hashed versions
// go worker(channel) {
// if true push to channel needs 1 buffer
// }

// Main function
// Create a password Channel
// For each password send it to the channel
// For loop that will create X number of worker functions
// Create a result Channel
//
