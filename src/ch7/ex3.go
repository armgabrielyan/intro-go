package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func totalArea(shapes ...Shape) float64 {
	var total float64

	for _, shape := range shapes {
		total += shape.area()
	}

	return total
}

func totalPerimeter(shapes ...Shape) float64 {
	var total float64

	for _, shape := range shapes {
		total += shape.perimeter()
	}

	return total
}

func main() {
	c := &Circle{radius: 3.5}
	r := &Rectangle{width: 2.67, height: 7.03}

	fmt.Printf("Area of the circle is %f\n", c.area())
	fmt.Printf("Perimeter of the circle is %f\n", c.perimeter())

	fmt.Printf("Area of the rectangle is %f\n", r.area())
	fmt.Printf("Perimeter of the rectangle is %f\n", r.perimeter())

	fmt.Printf("Total area is %f\n", totalArea(c, r))
	fmt.Printf("Total perimeter is %f\n", totalPerimeter(c, r))
}
