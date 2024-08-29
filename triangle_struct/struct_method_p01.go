package main

import (
	"fmt"
)

type Triangle struct {
	height float64
	bottom float64
}

func (t *Triangle) Area() float64 {
	return t.height * t.bottom / 2
}

func main() {
	tr := Triangle{3, 4}
	fmt.Println(tr.Area())
}
