package main

import "fmt"

func main() {
	var local1, local2, local3 int
	local1 = 10
	local2 = 20
	local3 = local1 + local2

	fmt.Printf("local1 = %d, local2 = %d, local3 = %d\n", local1, local2, local3)
}
