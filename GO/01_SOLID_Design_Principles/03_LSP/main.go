package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := &Square{}
	sq.width = size
	sq.height = size
	return sq
}

func (sq *Square) SetWidth(width int) {
	sq.width = width
	sq.height = width
}

func (sq *Square) SetHeight(height int) {
	sq.height = height
	sq.width = height
}

func UseIt(sized Sized) {
	w := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := w * 10
	actualArea := sized.GetHeight() * sized.GetWidth()

	fmt.Printf("expected area: %+v, actual area: %+v\n", expectedArea, actualArea)
}

func main() {
	rect := &Rectangle{
		width:  2,
		height: 3,
	}
	UseIt(rect)

	sq := NewSquare(5)
	UseIt(sq)
}
