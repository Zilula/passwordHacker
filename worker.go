package main
import "fmt"

func Worker(channel chan string, password string, listener chan string) {
	for pw := range channel{
		hashedPassword, _ := HashPassword(pw)
		fmt.Println(pw)
		if CheckPasswordHash("taylormade", hashedPassword) {
			listener <- "This hash matches the password: " + pw
		}
	} 

}

