package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{
		filename: filename,
	}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

// requirement: lazy loading of image till it needs to be Drawn. This can be done using virtual proxy

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{
		filename: filename,
		// the bitmap will only be initialised when it needs to be drawn
	}
}
func (lb *LazyBitmap) Draw() {
	if lb.bitmap == nil {
		lb.bitmap = NewBitmap(lb.filename)
	}
	fmt.Println("Drawing image", lb.filename)
}

// The reason why it is virtual is because when you create a lazy bitmap using the NewLazyBitmap, it hasn't been initialised yet, meaning theat the underlying implementation of the bitmap hasn't even been construted and it's only being constructed whenever somebody explicilty asks for it and in this case the whole thing gets constructed and subsquently used behind the scenes

func main() {
	bmp := NewBitmap("demo.png")
	DrawImage(bmp)
	fmt.Println("-------------------------")
	lbmp := NewLazyBitmap("demo2.png")
	DrawImage(lbmp)
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	_ = NewBitmap("demo.png")
	// DrawImage(bmp)
	fmt.Println("-------------------------")
	_ = NewLazyBitmap("demo2.png")
	// DrawImage(lbmp)
}
