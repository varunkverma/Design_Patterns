package main

import "fmt"

type Person struct {
	// address
	StreetName, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

// starting point for building up a Person
type PersonBuilder struct {
	person *Person // pointer to the person being built up
}

// Initialising this pointer of Person and PersonBuilder types
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		person: &Person{},
	}
}

// additional builders for building separate aspects of a Person Type

// All the specific builders embed the PersonBuilder, so, in a sense , they are also have access PersonBuilder feature, i.e., they are also of PersonBuilder type

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// we need to access specific builders via main builder, so we need to also initialise Specific builder via main builder. Again, since, these builders are in a way a type of the main builder, this was we have a way to transition from one to another

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{
		PersonBuilder: *pb,
	}
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{
		PersonBuilder: *pb,
	}
}

// we define fluent utility methods on specific builders

func (pab *PersonAddressBuilder) At(streetName string) *PersonAddressBuilder {

	pab.person.StreetName = streetName

	return pab
}

func (pab *PersonAddressBuilder) In(cityName string) *PersonAddressBuilder {

	pab.person.StreetName = cityName

	return pab
}

func (pab *PersonAddressBuilder) HasPostcode(postCode string) *PersonAddressBuilder {

	pab.person.Postcode = postCode

	return pab
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {

	pjb.person.CompanyName = companyName

	return pjb
}

func (pjb *PersonJobBuilder) AsA(positionName string) *PersonJobBuilder {

	pjb.person.Position = positionName

	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {

	pjb.person.AnnualIncome = annualIncome

	return pjb
}

// last point get the required built type
func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

// This way we have build a Domain Specific Languague for building up a person's information

func main() {
	pb := NewPersonBuilder()

	pb.
		Lives().
		At("123 London Road").
		In("London").
		HasPostcode("SW12ER1").
		Works().
		At("Microsoft").
		AsA("Senior Software Engineer").
		Earning(125000)

	person := pb.Build()

	fmt.Println(person)
}
