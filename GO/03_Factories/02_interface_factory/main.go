package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hello my name is %s, I'm %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Hello my name is %s, I'm lazy, I'm %d years old and a bit tired\n", p.name, p.age)
}

// the person's reference can be stored in the interface type that it implements
// This approach makes sure that you can only access the functions exposed by the interface and cannot access the underlying type
func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{
			name: name,
			age:  age,
		}
	}

	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p := NewPerson("John", 105)
	p.SayHello()
}
