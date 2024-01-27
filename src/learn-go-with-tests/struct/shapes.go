package main

import "math"

type Shape interface {
	Area() float64
}

// 四角形
type Rectangle struct {
	Width float64
	Height float64
}

// 四角形の周長
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// 四角形の面積
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 円形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
