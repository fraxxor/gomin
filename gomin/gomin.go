package main

import (
	"fmt"
	"de.fraxxor.gofrax/gomin/input/filereader"
)

func main() {
	fmt.Println("Hello World")
	content, err := filereader.ContentOf("D:/Programmierung/Go/rsrc/testfile.txt")
	if err != nil {
		fmt.Printf("Error on Read: %v", err)
	} else {
		fmt.Printf("Successfully read the following: %s", content)
	}
}