package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

/* Use struct as params type
 */
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

/* Use struct as receiver
 */
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

func main() {

	// Initialize variable with type Circle
	c := Circle{0, 0, 7}
	// field can be accessed using `.` operator.
	fmt.Println("Diameter of circle:", c.r*2)

	// circle area
	fmt.Println("Area of circle:", circleArea(c))

	rect := Rectangle{0, 0, 10, 10}
	fmt.Println("Area of Rectangle:", rect.area())
}
