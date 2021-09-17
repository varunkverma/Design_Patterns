package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

// Prototypes
var mainOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		StreetAddress: "123 East Dr",
		City:          "London",
	},
}

var auxiliaryOffice = Employee{
	Name: "",
	Office: Address{
		Suite:         0,
		StreetAddress: "66 West Dr",
		City:          "London",
	},
}

func (emp *Employee) DeepCopy() *Employee {
	// buffer
	b := &bytes.Buffer{}
	// encoder
	e := gob.NewEncoder(b)
	// encode employee's info
	_ = e.Encode(emp)

	// create a result emp Type
	newEmp := &Employee{}

	// create a decoder
	d := gob.NewDecoder(b)

	// decode store info into the new object
	_ = d.Decode(newEmp)

	return newEmp
}

// utility func
func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

// factory funcs
func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(
		&mainOffice, name, suite,
	)
}

func NewAuxiliaryOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(
		&auxiliaryOffice, name, suite,
	)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxiliaryOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
