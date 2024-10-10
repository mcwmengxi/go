package array

import (
	"fmt"
	"reflect"
	"testing"
)

// func main() {
// 	demo_1()
// 	fmt.Println("------")
// 	demo_2()
// 	fmt.Println("------")
// 	demo_3()
// }

func Ref() {
	x := 1
	p := &x
	*p = 2 
	fmt.Println(x,*p) // 输出 2

	str1 := "Golang"
	str2 := "Go语言"
	runeArr := []rune(str2)
	fmt.Println(reflect.TypeOf(str1).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2)：", len(str2))       // len(str2)： 8
	fmt.Println("runeArr:", string(runeArr[2]), len(runeArr))
}
func TestTimeAfter(t *testing.T) {
	Ref()
}

// 常量声明
// func demo_1() {
// 	const a uint8 = 10
// 	const str string = "1"
// 	const str_1, str_2 string = "str_1", "str_2"
// 	const str_3, str_4 = "str_33", "str_44"
// 	fmt.Println(a, str, str_1, str_2, str_3, str_4)
// }

// //变量声明
// func demo_2() {
// 	var num_1, num_2 uint8 = 1, 2
// 	num_3, num_4 := 3, 4

// 	name, age := "ymy", 5
// 	fmt.Println(num_1, num_2, num_3, num_4, name, age)

// 	var age_1 uint8 = 31
// 	var age_2 = 32
// 	age_3 := 33
// 	fmt.Println(age_1, age_2, age_3)

// 	var age_4, age_5, age_6 int = 31, 32, 33
// 	fmt.Println(age_4, age_5, age_6)

// 	var name_1, age_7 = "Tom", 30
// 	fmt.Println(name_1, age_7)

// 	name_2, is_boy, height := "Jay", true, 180.66
// 	fmt.Println(name_2, is_boy, height)
// }

// func demo_3() {
// 	fmt.Print("输出到控制台不换行")
// 	fmt.Println("输出到控制台并换行")
// 	fmt.Printf("name=%s,age=%d\n", "Tom", 30)
// 	fmt.Printf("name=%s,age=%d,height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
// }
