package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional factory
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

// structural factory
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		Name:         name,
		Position:     f.Position,
		AnnualIncome: f.AnnualIncome,
	}
}

func main() {
	// These are funcs retuned by the functional factory functions
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 100000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(developer, "\n", manager)

	// structural factoory usage
	bossFactory := NewEmployeeFactory2("CEO", 500000)
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)

	bossFactory.AnnualIncome += 10000
	boss2 := bossFactory.Create("Umora")
	fmt.Println(boss2)
}
