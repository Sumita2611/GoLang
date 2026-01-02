package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	//create a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	//add content to the file
	content := "Hello, this is a sample file created using Go!"
	_, error := io.WriteString(file, content+"\n")
	if error != nil {
		fmt.Println("Error writing to file:", error)
		return
	}

	fmt.Println("File created successfully")
}