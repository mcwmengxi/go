package slice

import "fmt"

func main() {
	slice_1()
	fmt.Println("------")
	slice_2()
	fmt.Println("------追加切片")
	slice_3()
	fmt.Println("------删除切片")
	slice_4()
}

// 声明切片
func slice_1() {
	var sli_1 = []int{1, 2, 3}
	fmt.Println("sli_1", sli_1, sli_1[len(sli_1)-1], cap(sli_1))
	var sli_2 = []int{} // 空切片
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli_2), cap(sli_2), sli_2)
	var sli_3 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)
	sli_4 := []int{1, 2, 3}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)
	var sli_5 []int = make([]int, 2, 3)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_5), cap(sli_5), sli_5)
	sli_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_6), cap(sli_6), sli_6)
}

// 截取切片
/*
	支持三种格式：
	无分号,取下标
	有一个分号，截取，左闭右开
	两个分号，前面规则不变，第二个分号后是cap截取的长度
*/

func slice_2() {
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice[1]=", sli[1], sli[:], sli[1:], sli[:4])
	fmt.Println("slice[0:3:4]=", sli[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3:4]), cap(sli[0:3:4]), sli[0:3:4])
}

//追加切片
func slice_3() {
	sli := []int{1, 3, 5}

	sli = append(sli, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)
	sli = append(sli, 9, 11)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)
}

// 删除切片
// 通过截取和追加方法实现删除效果
func slice_4() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//删除尾部 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[:len(sli)-2]), cap(sli[:len(sli)-2]), sli[:len(sli)-2])

	//删除开头 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[2:]), cap(sli[2:]), sli[2:])

	//删除中间 2 个元素
	sli = append(sli[:3], sli[3+2:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)
}
