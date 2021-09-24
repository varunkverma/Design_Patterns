package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

type PersonNameIterator struct {
	person  *Person
	current int // part of the name
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{
		person:  person,
		current: -1, // when you start iterating, you increated the value
	}
}

func (pni *PersonNameIterator) MoveNext() bool {
	pni.current++
	return pni.current < 3
}

func (pni *PersonNameIterator) Value() string {
	switch pni.current {
	case 0:
		return pni.person.FirstName
	case 1:
		return pni.person.MiddleName
	case 2:
		return pni.person.LastName
	}
	panic("We shouldn't be here")
}

func main() {
	p := &Person{
		FirstName:  "Alexander",
		MiddleName: "Graham",
		LastName:   "Bell",
	}
	for it := NewPersonNameIterator(p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
