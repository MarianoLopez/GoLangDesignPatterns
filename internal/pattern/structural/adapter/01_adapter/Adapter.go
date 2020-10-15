package main

import (
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1 int
	X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func newRectangle(width, height int) *VectorImage { //given this interface
	width -= 1
	height -=1
	return &VectorImage{[]Line {
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height}}}
}

//interface we have
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
		if pixel.X > maxX { maxX = pixel.X }
		if pixel.Y > maxY { maxY = pixel.Y }
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] { data[i][j] = ' ' }
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

func main() {
	rc := newRectangle(6, 4)
	a := VectorToRaster(rc)
	fmt.Println(DrawPoints(a))
}
//solution
type vectorToRasterAdapter struct {
	points []Point
}

func (receiver vectorToRasterAdapter) GetPoints()  []Point {
	return receiver.points
}

func (receiver *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			receiver.points = append(receiver.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			receiver.points = append(receiver.points, Point{x, top})
		}
	}

	fmt.Println("generated", len(receiver.points), "points")
}

func minmax(a, b int)  (int, int){
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return adapter
}
