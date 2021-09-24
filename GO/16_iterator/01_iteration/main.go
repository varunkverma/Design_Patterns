package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{
		p.FirstName,
		p.MiddleName,
		p.LastName,
	}
}

func main() {
	p := Person{
		FirstName:  "Alexander",
		MiddleName: "Graham",
		LastName:   "Bell",
	}

	for _, name := range p.Names() {
		fmt.Println(name)
	}
}
