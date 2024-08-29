package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	press   string
}

func main() {
	var bookGo Books

	bookGo.title = "Go 语言"
	bookGo.author = "www.runoob.com"
	bookGo.subject = "Go 语言教程"
	bookGo.press = "runoob"

	fmt.Println(bookGo.title)
	fmt.Println(bookGo.author)
	fmt.Println(bookGo.subject)
	fmt.Println(bookGo.press)
}
