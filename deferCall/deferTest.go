package main

import "fmt"

func deferFunc() {
	fmt.Println("deferFunc up")
	defer fmt.Println("deferFunc 01")
	defer func() {
		fmt.Println("deferFunc 02 and closure")
	}()
	defer func() {
		defer func01()
		defer func02()
		defer func03()
	}()
}

func func01() {
	fmt.Println("func01")
}
func func02() {
	fmt.Println("func02")
}
func func03() {
	fmt.Println("func03")
}

func main() {
	deferFunc()
	// func01()
	// func02()
	// func03()
}
