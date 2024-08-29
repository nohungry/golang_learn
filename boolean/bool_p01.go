package main

import "fmt"

func boolTest() {
	var a bool = true
	var b bool = false
	fmt.Println("a:", a, "b:", b)
}

func main() {
	boolTest()
	varTest()
}

func varTest() {
	var aVer int = 100
	fmt.Println("aVer:", aVer)
	fmt.Printf("aVer type: %T\n", aVer)
	fmt.Println(aVer == 50)
}
