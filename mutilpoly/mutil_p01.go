package main

import (
	oop1 "mut/o1"
	oop2 "mut/o2"
)

// type Num int

// func (x Num) equal(y Num) bool {
// 	return x == y
// }

// type NumI interface {
// 	equal(y Num) bool
// }

// var x Num = 8
// var y NumI = &x

type Num int

func (x Num) Equal(y int) bool {
	return int(x) == y
}

func (x Num) LessThan(y int) bool {
	return int(x) < y
}

func (x Num) BiggerThan(y int) bool {
	return int(x) > y
}

func main() {
	var f1 Num = 8
	var f2 oop1.NumInterface1 = f1
	var f3 oop2.NumInterface2 = f2
}
