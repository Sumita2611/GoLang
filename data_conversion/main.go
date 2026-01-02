package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("number is ", num)
	fmt.Printf("Type of num is %T\n", num)

	// num = num + 1.23 // this will give error as we are trying to add float to int

	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("data is ", data)
	fmt.Printf("Type of data is %T\n", data)

	num = 123
	str := strconv.Itoa(num) // int to string conversion
	fmt.Println("str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "456"
	number_int, _ := strconv.Atoi(number_string) // string to int conversion
	fmt.Println("number_int is ", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64) // string to float conversion
	fmt.Println("number_float is ", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)
}