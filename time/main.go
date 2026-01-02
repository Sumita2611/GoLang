package main

import (
	"fmt"
	"time"
)

func main() {	
currTime := time.Now()
	fmt.Println("Current Time: ", currTime)
	fmt.Printf("Type of currTime is %T\n", currTime)

	formatted := currTime.Format("02-01-2006, 3:04 PM");
	fmt.Println("Formatted Time: ", formatted)

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted Time from string: ", formatted_time)

	//add one more day in curr time
	new_date := currTime.Add(24 * time.Hour)
	fmt.Println("New Date after adding one day: ", new_date)
	formatted_new_date := new_date.Format("02-01-2006, Monday");
	fmt.Println("Formatted New Date: ", formatted_new_date)
}