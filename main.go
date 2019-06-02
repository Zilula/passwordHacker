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
	for i := 0; i < 10; i++ {
		go Worker(passwordChannel, hashedPasswordTest, listenerChannel)
	}

	// Sends a password to the channel
	for _, password := range passwordSlice {
		passwordChannel <- password
	}
	fmt.Println(<-listenerChannel)
}
