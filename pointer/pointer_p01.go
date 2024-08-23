package main

import "fmt"

func pointerTest01() {
	var score int = 100
	var name string = "Barry"
	fmt.Printf("%p\n", &score)
	fmt.Printf("%p\n", &name)
}

func main() {
	pointerTest01()
}
