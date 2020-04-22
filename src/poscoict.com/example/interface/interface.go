package main

import "strconv"

// 인터페이스 정의
type Shape interface {
	area() int // 넓이
	show() string
}

// Rect 구조체 정의
type Rect struct {
	width  int
	height int
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() int {
	return r.width * r.height
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) show() string {
	return "width : " + strconv.Itoa(r.width) + ", height : " + strconv.Itoa(r.height)
}

func main() {
	r := Rect{10, 20}
	println(r.area())
	println(r.show())
}
