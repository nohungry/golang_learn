package go_test001

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	array := []int{6, 7, 9}
	ret := Min(array)
	fmt.Println(ret)
}
