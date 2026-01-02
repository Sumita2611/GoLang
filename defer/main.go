package main
import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

//when two defers : follows LIFO principle
func main() {
	fmt.Println("Starting of the program")
	data := add(3, 4)
	defer fmt.Println("Data is : ", data)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}