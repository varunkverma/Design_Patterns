package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square of size %f", s.Side)
}

// Decorator - but not a good approach as we would need to create one struct for every specific one, so
// type ColoredSquare struct{
// 	Square
// 	Color string
// }

// Better to use an interface
// Decorator
type ColoredShape struct {
	Shape Shape // interface
	Color string
}

func (cs *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", cs.Shape.Render(), cs.Color)
}

// Decorator
type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (ts *TransparentShape) Render() string {
	return fmt.Sprintf("%s has the transparency of %d%%", ts.Shape.Render(), int(ts.Transparency*100.0))
}

func main() {
	circle := &Circle{
		Radius: 2,
	}

	fmt.Println(circle.Render())

	redCircle := &ColoredShape{
		Shape: circle,
		Color: "Red",
	}

	fmt.Println(redCircle.Render())

	// The drawback is once you have packed a particular shape in the decorator, you cannot access the methods particularly associated to the internal type and not mentioned in the interface
	// redCircle.Resize() // error

	// But one advantage is that Decorators can be composed, i.e., you can apply decorators to decorators

	redHalfTransparent := &TransparentShape{
		Shape:        redCircle,
		Transparency: 0.5,
	}
	fmt.Println(redHalfTransparent.Render())
}
