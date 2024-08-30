package main

import "fmt"

func hashtableCreate(array []int16) (hashTable map[int16]int16) {
	hashTable = make(map[int16]int16)
	for i := 0; i < len(array); i++ {
		hashTable[array[i]] = int16(i)
	}
	return hashTable
}

func main() {
	target := int16(21)
	var array = []int16{63, 21, 34, 74, 85}
	mapTest := hashtableCreate(array)
	// fmt.Println(mapTest)
	res := matchNum(mapTest, target)
	fmt.Println(res)
}

func matchNum(hashTable map[int16]int16, target int16) (result map[int16]int16) {
	for key, index := range hashTable {
		if key == target {
			result = map[int16]int16{index: key}
		} else {
			// handle the case when the key does not match the target
			result = map[int16]int16{}
		}
	}
	return result
}
