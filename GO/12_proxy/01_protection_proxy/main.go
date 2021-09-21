package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    *Car
	driver *Driver
}

func NewCarProxy(d *Driver) *CarProxy {
	return &CarProxy{
		car:    &Car{},
		driver: d,
	}
}

func (cp *CarProxy) Drive() {
	if cp.driver.Age >= 16 {
		cp.car.Drive()
	} else {
		fmt.Println("driver too young to drive")
	}
}

func main() {
	car := NewCarProxy(&Driver{Age: 12})
	car.Drive()

	car.driver.Age = 22
	car.Drive()
}
