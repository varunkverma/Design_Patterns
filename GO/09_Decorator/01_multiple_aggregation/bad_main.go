package main

// import "fmt"

// type Bird struct {
// 	Age int
// }

// func (b *Bird) Fly() {
// 	if b.Age >= 10 {
// 		fmt.Println("Flying")
// 	}
// }

// type Lizard struct {
// 	Age int
// }

// func (l *Lizard) Crawl() {
// 	if l.Age < 10 {
// 		fmt.Println("Crawling")
// 	}
// }

// func (d *Dragon) GetAge() int {
// 	return d.Bird.Age
// }

// func (d *Dragon) SetAge(age int) {
// 	d.Bird.Age = age
// 	d.Lizard.Age = age
// }

// // issue as both Bird and Lizard types have a similar field called Age
// type Dragon struct {
// 	Bird
// 	Lizard
// }

// func main() {
// 	// 	d := &Dragon{}
// 	// 	// d.Age=1 // ambiguous
// 	// 	// d.Bird.Age = 11 // workable but inconvinient
// 	// 	// d.Lizard.Age = 11

// 	// 	d.SetAge(11) // but inter Age field is possible to be accessed.

// 	// 	d.Fly()
// }
