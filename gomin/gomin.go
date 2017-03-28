package main

import (
	"fmt"
	"de.fraxxor.gofrax/gomin/input/filereader"
)

func main() {
	fmt.Println("Hello World")
	reader := new(filereader.Filereader)
	content, err := reader.ContentOf("D:/Programmierung/Go/rsrc/testfile.txt")
	if err != nil {
		fmt.Printf("Error on Read: %v", err)
	} else {
		fmt.Printf("Successfully read the following:\n%s", content)
	}
}