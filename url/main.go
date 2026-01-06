package main 

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learning URL")
	myURL := "sumita"
	fmt.Printf("Type of url : %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Printf("Type of parsedURL : %T\n", parsedURL)

}