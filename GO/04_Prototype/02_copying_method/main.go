package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p // we are not copying the reference, just the value
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	john := Person{
		Name: "John",
		Address: &Address{
			StreetAddress: "123 lon street",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Meg"},
	}

	jane := john.DeepCopy() // here the Person object is copied into jane

	jane.Name = "Jane" // okay
	jane.Address.StreetAddress = "144 Kolls street"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
