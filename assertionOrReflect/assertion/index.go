package assertion

import "fmt"
type Animial struct {
	Name string
}
type Person struct {
	Name string
}

func (p Person) create(val string){
	fmt.Println(val, "无限创造" )
} 
func AssertionFunc(){
	// 断言Assertion
	// p := Person{
	// 	Name: "人类",
	// }
	// is(p)
	//断言 把一个接口类型指定为它原始的类型
	is(Animial{})

}
func is(v interface{}){
	switch v.(type){

	case Animial:
		fmt.Println("动物")
	case Person:
		v.(Person).create(v.(Person).Name)
	}
}