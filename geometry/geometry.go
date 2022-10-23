package geometry

import "math"

type Rectangle struct {
	width  float64
	height float64
}
type Triangle struct{
	height,length float64

}
type Circle struct {
	radius float64
}
type Shape interface{
	Area() float64
}

func Perimeter(rec Rectangle) float64 {
	perimeter := 2 * (rec.width + rec.height)
	return perimeter
}

func (rec Rectangle) Area() float64 {
	area := rec.width * rec.height
	return area
}

func (cir Circle) Area()float64 {
	area := math.Pi * (cir.radius*cir.radius)
	return area
}

func (tri Triangle) Area() float64 {
	return (tri.height*tri.length)/2
}