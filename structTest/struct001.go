package main

import "fmt"

func structTest() {
	type person struct {
		name string
		age  int
	}

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}

func main() {
	structTest()
}

func temp() {
	var a []int
	a = append(a, 1)
}
