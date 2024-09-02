package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	var num int
// 	typ := reflect.TypeOf(num)
// 	val := reflect.ValueOf(num)
// 	zeroVal := reflect.Zero(typ)

// 	fmt.Printf("Type: %v\n", typ)           // Type: int
// 	fmt.Printf("Value: %v\n", val)          // Value: 0
// 	fmt.Printf("Zero Value: %v\n", zeroVal) // Zero Value: 0
// }

// func main() {
// 	// 創建一個整數變數
// 	var num int = 42

// 	// 使用 reflect.ValueOf 取得 reflect.Value
// 	valueOfNum := reflect.ValueOf(&num).Elem()

// 	// 使用 Elem() 取得可寫的value
// 	if valueOfNum.CanSet() {
// 		// 使用 SetInt 設定新的整數value
// 		valueOfNum.SetInt(99)
// 	}

// 	fmt.Println("Updated Value:", num) // 99
// }

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}
func main() {
	calculator := Calculator{}

	// 使用 reflect.ValueOf 取得計算器物件的reflect value
	val := reflect.ValueOf(calculator)

	// 使用 MethodByName 方法取得名為 "Add" 的方法的reflect value
	method := val.MethodByName("Add")

	// 準備方法的參數
	args := []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(3)}

	// 調用方法並取得結果，然後轉換為整數型別
	result := method.Call(args)[0].Interface()

	fmt.Printf("Result: %d\n", result) // Result: 8
}
