package reflect

import (
	"fmt"
	"reflect"
)

type Animial struct {
	Name string
	Owner string
}
type Person struct {
	Name string
	tag string
	Animial
}

func (p Person) Create(val string){
	fmt.Println(val, "无限创造" )
} 
func ReflectFunc(){
	ani := Animial{
		Name: "dog",
		Owner: "empty",
	}
	p := Person{
		Name: "人类",
		tag: "yys",
		Animial: ani,
	}
	is(p)
// 	普通反射

// struct反射
// 匿名或嵌入字段的反射
// 判断传入的类型是否是我们想要的类型
// 通过反射修改内容
// 通过反射调用方法
}

func is(v interface{}){
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	fmt.Println("reflect.typeOf", t,val)
	// for i:=0;i< t.NumField();i++  {
	// 	fmt.Println(val.Field(i))
	// }
	// fmt.Println("reflect.ValueOf.FieldByName", val.FieldByName("Name"),val.FieldByName("Owner"))
	// p := val.FieldByIndex([]int{0,1})
	// fmt.Println("reflect.ValueOf.FieldByIndex",val.FieldByIndex([]int {2, 0}))

	// typ := t.Kind()
	// if typ == reflect.Struct {
	// 	fmt.Println("Struct", typ)
	// }
	// if typ == reflect.String {
	// 	fmt.Println("String", typ)
	// }
	// 修改原始值 传入的v是一个地址才能修改
	// origin := val.Elem()
	// origin.FieldByName("Name").SetString("人类plus")
	// fmt.Println(v,origin)

	// 方法需要大写，使之不再是私有化的
	methed := val.Method(0)
	methed.Call([]reflect.Value{reflect.ValueOf("沙盒游戏")})
}