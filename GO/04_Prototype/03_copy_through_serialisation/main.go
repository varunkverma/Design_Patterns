package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	// create a buffer
	b := &bytes.Buffer{}

	// create an encoder
	e := gob.NewEncoder(b)

	// encode the type's object into the buffer
	_ = e.Encode(p)

	// fmt.Println(b.String())

	// create a decoder to read from buffer
	d := gob.NewDecoder(b)

	result := &Person{} // the new object

	// decode the bytes into the passed reference of an object
	_ = d.Decode(result)

	return result
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
