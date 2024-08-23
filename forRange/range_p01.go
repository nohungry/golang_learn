package main

import "fmt"

// func main() {
// 	// numbers := []int{2, 3, 4, 5, 6}

// 	// for i, number := range numbers {
// 	// 	fmt.Println("Index:", i, "Number:", number)
// 	// }

// 	var i int
// JumpLoop:
// 	for ; ; i++ {
// 		if i > 10 {
// 			break JumpLoop
// 		}
// 		fmt.Println(i)
// 	}
// }

func main() {
	// for key, value := range []int{11, 32, 53, 24, 55} {
	// 	fmt.Println(key, value)
	// }
	// mapRange()
	// sliceRange()
	channelRange()
}

func mapRange() {
	m := map[int]string{
		1: "A",
		2: "BB",
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func sliceRange() {
	s := []int{1, 2, 3, 4, 5}
	for key, value := range s {
		fmt.Println(key, value)
	}
}

func channelRange() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
