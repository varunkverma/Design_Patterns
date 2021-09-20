package main

import "fmt"

type Buffer struct {
	width, height int
	buffer        []rune
}

// factory
func NewBuffer(width int, height int) *Buffer {
	return &Buffer{
		width:  width,
		height: height,
		buffer: make([]rune, width*height),
	}
}

func (bf *Buffer) At(index int) rune {
	return bf.buffer[index]
}

// Viewport shows just a particular part of the buffer.
type Viewport struct {
	buffer *Buffer
	offset int
}

// factory
func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{
		buffer: buffer,
	}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// and such there could be many other types of viewports, so in order to access them with ease, we need a facade to incorporate information about different buffers and viewports

// facade
type Console struct {
	buffer    []*Buffer
	viewports []*Viewport
	offset    int
}

// factory
func NewConsole() *Console {
	// default Console
	b := NewBuffer(200, 150)
	v := NewViewport(b)

	return &Console{
		buffer:    []*Buffer{b},
		viewports: []*Viewport{v},
		offset:    0,
	}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	// interacting with only facade
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
