package main

import "fmt"

func Worker(channel chan string, password string, listener chan string) {
	for pw := range channel {
		hashedPassword, _ := HashPassword(pw)
		fmt.Println(pw)
		if CheckPasswordHash("taylormade", hashedPassword) {
			listener <- "This hash matches the password: " + pw
			// return "This hashed matches the password: " + pw
			return
		} 
		// else {
		// 	listener <- "This hash did not match any passwords: " + pw
			
		// }
	// }
	// return "did not match anything"
	}
}