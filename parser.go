package main

import (  
    "fmt"
	"io/ioutil"
	"strings"
	// "reflect"
)

func Parser() []string {  
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return nil
	}

	// fmt.Println("Contents of file:", string(data))
	// fmt.Println("type of", reflect.TypeOf(string(data)))
	
	txtString := string(data)
	
	parsedTxt := strings.Split(txtString, "\n")
	
	// fmt.Println("parsed", parsedTxt)
	// fmt.Println("parsed", parsedTxt[2])
	// fmt.Println("type of", reflect.TypeOf(parsedTxt))

	return parsedTxt
	
}