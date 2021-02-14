package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Around() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Around() float64 {
	return 2*r.length + 2*r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Around() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rectangle := Rectangle{10, 3}
	circle := Circle{10}
	calculate(rectangle, "Rectangle")
	calculate(circle, "Circle")

}

func calculate(s Shape, shape string) {
	fmt.Println(shape+" Area : ", s.Area())
	fmt.Println(shape+" Around : ", s.Around())
}
