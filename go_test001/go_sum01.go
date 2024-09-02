package go_test001

func Min(arr []int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
