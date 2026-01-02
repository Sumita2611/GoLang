package main

import "fmt"

func main() {
	x := 18

	if x == 18 {
		fmt.Println("You are eligible to vote!")
	} else if x > 18 {
		fmt.Println("You are above 18 years!")
	} else {
		fmt.Println("You are not eligible to vote!")
	}
}
