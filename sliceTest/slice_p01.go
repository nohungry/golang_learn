package main

import "fmt"

func main() {
	// slice of int
	var s []int
	// fmt.Println(s)
	for i := 0; i < 100; i += 2 {
		s = append(s, i)
	}
	fmt.Println(s)

	// slice of s
	s1 := s[1:5]
	fmt.Println(s1)
}
