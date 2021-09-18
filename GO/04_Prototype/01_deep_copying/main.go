package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{
		Name: "John",
		Address: &Address{
			StreetAddress: "123 lon street",
			City:          "London",
			Country:       "UK",
		},
	}

	jane := john // here the Person object is copied into jane

	jane.Name = "Jane" // okay

	// assign jane's address pointer a new Address type reference
	jane.Address = &Address{
		StreetAddress: john.Address.StreetAddress,
		City:          john.Address.City,
		Country:       john.Address.Country,
	}
	jane.Address.StreetAddress = "144 Kolls street"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
