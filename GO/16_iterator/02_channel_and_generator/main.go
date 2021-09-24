package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		out <- p.FirstName

		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}

		out <- p.LastName
	}()
	return out
}

func main() {
	p := Person{
		FirstName:  "Alexander",
		MiddleName: "Graham",
		LastName:   "Bell",
	}

	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}
}
