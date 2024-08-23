package main

import "fmt"

func switchCase() int {
	var a = "love"
	switch a {
	case "love":
		fmt.Println("I love you")
	case "programming":
		fmt.Println("I love programming")
	default:
		fmt.Println("I love nothing")
	}
	return 0
}

func main() {
	switchCase()
}
