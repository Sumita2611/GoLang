package main

import "fmt"

func main() {
	fmt.Println("Learning array in Go!")
	var arr [5]string
	arr[0] = "Apple"
	arr[1] = "Banana"
	
	fmt.Println("Fruit array:", arr)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array:", numbers)

	fmt.Println("Length of numbers array:", len(numbers))

	var price[5]string
	fmt.Println("Price is:", price)
	fmt.Printf("Price Array is %q\n", price)
}