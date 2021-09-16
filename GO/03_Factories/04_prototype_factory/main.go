package main

import "fmt"

type Employee struct {
	Name         string
	Position     string
	AnnualIncome int
}

type EmployeeRole int

const (
	DEVELOPER EmployeeRole = iota
	MANAGER
)

func NewEmployee(role EmployeeRole) *Employee {
	switch role {
	// pre-configured objects
	case DEVELOPER:
		return &Employee{"", "developer", 60000}
	case MANAGER:
		return &Employee{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {

	m := NewEmployee(MANAGER)
	m.Name = "Same"

	fmt.Println(m)
}
