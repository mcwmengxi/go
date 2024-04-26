package array

import "fmt"

func main() {
	array_1()
	fmt.Println("---------")
	array_2()
	fmt.Println("---------")
	transfer()
	fmt.Println("---------")
	transfer_1()
	fmt.Println("---------")
	transfer_2()
}

// 一维数组
func array_1() {
	var arr_1 = [5]int{1, 2, 3, 4, 5}
	arr_2 := [5]int{1, 2, 3, 4}
	arr_3 := [...]int{1, 4}

	// 不指定以最大下标为长度
	arr_4 := [...]int{9: 4, 0: 6, 1: 8}
	arr_5 := [11]int{9: 4, 0: 6, 1: 8}
	// 默认全为0
	var arr_6 [5]int
	fmt.Println(arr_1, arr_2, arr_3, arr_4, arr_5, arr_6)
}

// 二维数组
func array_2() {
	var arr_1 = [2][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	arr_2 := [2][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	// 不指定默认为0,只能任意外层数组的个数
	arr_3 := [...][3]int{{1, 2, 4}, {2, 4}, {1: 1, 0: 1, 2: 4}}

	// 指定数组长度后不可扩展
	arr_4 := [...]int{2: 4, 0: 6, 1: 8}
	arr_5 := arr_4
	fmt.Println(arr_1, arr_2, arr_3, arr_4, arr_5, arr_4 == arr_5)
}

func transfer() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr(arr)
	fmt.Println(arr)
}
func modifyArr(a [5]int) {
	a[1] = 20
	fmt.Println("qaq", a)
}

// &取地址符
func transfer_1() {
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArr_1(&arr)
	fmt.Println(arr, &arr)
}
func modifyArr_1(a *[5]int) {
	a[1] = 20
	fmt.Println("qaq_1", a, &a)
}

//数组赋值
func transfer_2() {
	var arr = [5]int{1, 2, 3, 4, 5}
	arr_1 := arr
	// var arr_2 [6] int = arr // cannot use arr (variable of type [5]int) as [6]int value
	fmt.Println(arr_1, arr_1 == arr)
}
