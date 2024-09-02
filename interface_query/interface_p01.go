// package main

// import "fmt"

// func main() {
// 	slice := make([]int, 0)
// 	slice = append(slice, 6, 7, 8)
// 	var I interface{} = slice
// 	if res, ok := I.([]int); ok {
// 		fmt.Println(res)
// 		fmt.Println("ok")
// 	}
// }

package main

import "fmt"

func main() {
	// 初始化一个包含切片的空接口
	slice := []int{6, 7, 8}
	// var I interface{} = slice
	// 如果需要同时存储多个不同类型的值，可以使用一个切片或 map
	mixedSlice := []interface{}{slice, 42, "hello", 3.14}
	for _, v := range mixedSlice {
		fmt.Println(v)
	}
}
