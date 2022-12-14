package main

import (
	"fmt"
	"math"
)

/*
--------------------------------Interface and Structs ----------------------------------------------------------------
*/

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.area()
	}
	return area
}

func totalSurfaceArea(shapes ...Shape) float64 {
	var surfaceArea float64
	for _, shape := range shapes {
		surfaceArea += shape.perimeter()
	}
	return surfaceArea
}

type Circle struct {
	x, y, r float64
}

/* when implementing these functions for the interface, you have to make sure the receiver is the struct, not the pointer to the struct. This means you should not implment the methods with the asterisks before the struct name, unlike the book instructs us to do. 
*/
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	l, w float64
}

func (r Rectangle) area() float64 {
	return r.l * r.w 
}

func (r Rectangle) perimeter() float64 {
	return 2 * r.l + 2 * r.w
}

/*
--------------------------------Main Function ----------------------------------------------------------------
 */
func main() {

	c1 := Circle{1,1,6}
	c2 := Circle{7, -7, 13}
	r1 := Rectangle{45, 23}
	r2 := Rectangle{ 13,25}

	fmt.Println(totalArea(c1,c2,r1,r2))
	fmt.Println(totalSurfaceArea(c1,c2,r1,r2))


}