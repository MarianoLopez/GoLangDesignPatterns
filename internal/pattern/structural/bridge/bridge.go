package main

import "fmt"

//avoid Cartesian product
//Circle, square
//Raster, vector
//=> RasterCircle, VectorCircle, RasterSquare

type Render interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {

}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius ", radius)
}

type RasterRender struct {
	Dpi int
}

func (r *RasterRender) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius ", radius)
}

type Circle struct {
	renderer Render
	radius float32
}

func newCircle(renderer Render, radius float32) *Circle {
	return &Circle{
		renderer: renderer,
		radius:   radius,
	}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize	(factor float32)  {
	c.radius *= factor
}

func main() {
	//raster := RasterRender{}
	vector := VectorRenderer{}
	circle := newCircle(&vector, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
}
