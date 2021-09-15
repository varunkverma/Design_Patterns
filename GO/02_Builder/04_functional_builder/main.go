package main

import "fmt"

type Person struct {
	name, profession string
}

type personMod func(*Person)

// Functional builder
type PersonBuilder struct {
	actions []personMod
}

// fluent methods assocaited to the Builder
func (pb *PersonBuilder) Called(name string) *PersonBuilder {
	addNameAction := func(p *Person) {
		p.name = name
	}
	pb.actions = append(pb.actions, addNameAction)
	return pb
}

func (pb *PersonBuilder) WorksAsA(profession string) *PersonBuilder {

	pb.actions = append(pb.actions, func(p *Person) {
		p.profession = profession
	})
	return pb
}

func (pb *PersonBuilder) Build() *Person {
	p := &Person{}

	// these functions are executed lazily
	for _, action := range pb.actions {
		action(p)
	}
	return p
}

func main() {

	pb := &PersonBuilder{}
	pb.
		Called("Xio").
		WorksAsA("Developer")

	p := pb.Build()

	fmt.Printf("%+v\n", p)
}
