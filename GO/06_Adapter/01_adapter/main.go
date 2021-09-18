package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1 int // starting point of a line
	X2, Y2 int // ending point of a line
}

type VectorImage struct {
	Lines []Line
}

//Provided method that avaible to create a rectangle, assuming a 3rd party func
func NewRectangle(width, height int) *VectorImage {
	width -= 1 // things starts from 0
	height -= 1
	return &VectorImage{
		Lines: []Line{
			{0, 0, width, 0},           //top
			{0, 0, 0, height},          // left
			{width, 0, width, height},  // right
			{0, height, width, height}, // bottom
		},
	}
}

// The interface We use in our code
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// adapter
type vectorToRasterAdapter struct {
	points []Point
}

// vectorToRasterAdapter is a type of RasterImage as it implements GetPoints() []Point function
func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

var pointCache = map[[16]byte][]Point{}

func (v *vectorToRasterAdapter) addLineCached(line Line) {

	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	lineHashValue := hash(line)
	if cachedPoints, ok := pointCache[lineHashValue]; ok {
		v.points = append(v.points, cachedPoints...)
		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}

	pointCache[lineHashValue] = v.points

	fmt.Println("generated", len(v.points), "points")
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}

	fmt.Println("generated", len(v.points), "points")
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLineCached(line)
		// adapter.addLine(line)
	}

	return adapter
}

func main() {
	rc := NewRectangle(6, 4)
	roasterImage := VectorToRaster(rc)
	_ = VectorToRaster(rc)
	fmt.Println(DrawPoints(roasterImage))
}
