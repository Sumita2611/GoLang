package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area string
	State string
}

//custom type
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var Ram Person
	fmt.Println("Person Details:", Ram)

	Ram.FirstName = "Ram"
	Ram.LastName = "Sharma"
	Ram.Age = 30
	fmt.Println("Updated Person Details:", Ram)

	// another method to create struct variable
	person1 := Person{
		FirstName: "Sita",
		LastName:  "Verma",
		Age:       28,
	}
	fmt.Println("Person1 Details:", person1)

	// with new keyword
	person2 := new(Person)
	person2.FirstName = "Gita"
	person2.LastName = "Kumari"
	person2.Age = 25
	fmt.Println("Person2 Details:", *person2)

	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Rohan",
		LastName:  "Singh",
		Age:       32,
	}
	employee1.Person_Contact = Contact{
		Email: "rohan.singh@example.com",
		Phone: "9876543210",
	}
	employee1.Person_Address = Address{
		House: 101,
		Area:  "Sector 12",
		State: "Delhi",
	}

	fmt.Println("Employee1 Details:", employee1)
}