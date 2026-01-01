package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func add1(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("Learning functions in Go!")
	ans := add(3, 4)
	fmt.Println("addition of two numbers:", ans)
}