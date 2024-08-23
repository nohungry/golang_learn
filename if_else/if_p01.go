package main

import "fmt"

func ExampleIfElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// Output: 7 is odd
}

func main() {
	ExampleIfElse()
}
