package main

import "fmt"

func test001() {
	for i := -1; i < 10; i++ {
		fmt.Println(i)
	}
}

func test002() {
	i := 0
	for {
		i++
		if i > 50 {
			break
		}
		fmt.Println(i)
	}
}

func main() {
	// test001()
	test002()
}
