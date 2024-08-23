package main

import "fmt"

func strTest01() string {
	str01 := "\"GO Web Programming\", I love this book.\n"
	str02 := `"Go WEB",
			I love this book.\n`

	fmt.Println(str01)
	fmt.Println(str02)

	return "str00000"
}

func main() {
	strTest01()
}
