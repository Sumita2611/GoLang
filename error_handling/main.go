package main

import "fmt"

func divide(a,b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	// fmt.Println("Started error handling in Go!")
	// ans, err := divide(10,0);
	// if err != nil {
	// 	fmt.Println("Error handling");
	// }
	// fmt.Println("Division of two numbers:", ans)

	ans, _ := divide(10,0);
	fmt.Println("Division of two numbers:", ans)
}
