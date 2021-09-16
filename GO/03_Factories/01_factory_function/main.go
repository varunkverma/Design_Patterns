package main

import "fmt"

type Person struct {
	Name          string
	Age, EyeCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name:     name,
		Age:      age,
		EyeCount: 2,
	}
}

func main() {
	p := NewPerson("John", 33)
	fmt.Println(p)
}
