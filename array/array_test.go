package array

import (
	"fmt"
	"testing"
)

func getArrByIndex(arr [3]int, i int) int {
	// 数组越界 通过defer recover捕获 类似与try catch
	defer func() {
		fmt.Println("123")
		if r := recover(); r != nil {
			fmt.Println(r)
			i = -1
		}
	}()
	fmt.Println(arr[i])
	return arr[i]
}
func TestArray(t *testing.T) {

	var arr = [3]int{1, 2, 3}
	getArrByIndex(arr, 5)
}