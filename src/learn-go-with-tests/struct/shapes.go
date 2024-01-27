package main

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
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
