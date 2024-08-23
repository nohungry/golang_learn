package main

import "fmt"

// const (
// 	e  = 2.71828182845904523536028747135266249775724709369995957496696763
// 	pi = 3.14159265358979323846264338327950288419716939937510582097494459
// )

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func main() {
	// fmt.Println("e:", e, "pi:", pi)
	fmt.Println("North:", North, "East:", East, "South:", South, "West:", West)
}
