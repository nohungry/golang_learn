package main

import "fmt"

func hashtableCreate(array []int) (hashTable map[int]int) {
	hashTable = make(map[int]int)
	for i := 0; i < len(array); i++ {
		hashTable[array[i]] = i
	}
	return hashTable
}

func main() {
	var array = []int{63, 21, 34, 74, 85}
	mapTest := hashtableCreate(array)
	fmt.Println(mapTest)
}
