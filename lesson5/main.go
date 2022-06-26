package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Type() string
	Area() int
}

type Rectangle struct {
	name   string
	width  int
	height int
}

func (r *Rectangle) Type() string {
	return r.name
}

func (r *Rectangle) Area() float64 {
	return float64(r.width * r.height)
}

type Circle struct {
	name   string
	radius int
}

func (c *Circle) Type() string {
	return c.name
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.radius*2)
}

func main() {
	rectangle := NewRectangle(30, 40)
	circle := NewCircle(10)

	fmt.Println(rectangle.Type())
	fmt.Println(rectangle.Area())

	fmt.Println(circle.Type())
	fmt.Println(circle.Area())
}

func NewRectangle(width int, height int) *Rectangle {
	return &Rectangle{
		"Rectangle",
		width,
		height,
	}
}

func NewCircle(radius int) *Circle {
	return &Circle{
		"Circle",
		radius,
	}
}
