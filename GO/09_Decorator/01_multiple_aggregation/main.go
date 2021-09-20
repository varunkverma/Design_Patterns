package main

import "fmt"

type Aged interface {
	GetAge() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) GetAge() int    { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) GetAge() int { return l.age }
func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling")
	}
}

// sort of a Decorator, as it provides additional functionalities to set ages in a consistent manner
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) GetAge() int { return d.bird.age }
func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

// factory func to create Dragon
func NewDragon() *Dragon {
	return &Dragon{
		bird:   Bird{},
		lizard: Lizard{},
	}
}

func main() {
	d := NewDragon()
	d.SetAge(10)
	d.Fly()

	d.SetAge(5)
	d.Crawl()
}
