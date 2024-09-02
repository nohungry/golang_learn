package main

import "fmt"

type Triangle struct {
	bottom float32
	height float32
}

type square struct {
	side float32
}

func (t *Triangle) Area() float32 {
	return 0.5 * t.bottom * t.height
}

func (s *square) Area() float32 {
	return s.side * s.side
}

type Shape interface {
	Area() float32
}

func main() {
	t := &Triangle{bottom: 10, height: 5}
	s := &square{side: 5}
	shape := []Shape{t, s}
	for n, _ := range shape {
		fmt.Println(shape[n])
		fmt.Println(shape[n].Area())
	}
}
