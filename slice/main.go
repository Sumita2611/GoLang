package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 2, 4, 6, 8, 3, 0)
	// fmt.Println("Slice of numbers:", numbers)
	// fmt.Printf("Numbers has data type: %T\n", numbers)
	// fmt.Println("Length of numbers:", len(numbers))

	//creating slice using make function : len and cap
	numbers := make([]int, 3, 5)
	numbers = append(numbers, 4)
	numbers = append(numbers, 41)
	numbers = append(numbers, 42)
	fmt.Println("Slice of numbers:", numbers)
	fmt.Println("Length of numbers:", len(numbers))
	fmt.Println("Capacity of numbers:", cap(numbers)) // capacity doubles when length exceeds capacity

	// create empty slice
	stock := []string{}
	fmt.Println("Stock slice:", stock)
	fmt.Println("Length of stock slice:", len(stock))
	fmt.Println("Capacity of stock slice:", cap(stock))
}