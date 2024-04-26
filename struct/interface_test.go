package structdemo

import (
	"fmt"
	"testing"
)

type Animate interface {
	getName() string
}

type Cat struct {
	Name string
}

func (c Cat) getName() string {
	return c.Name
}

type Dog struct {
	Name string
}

func (d Dog) getName() string {
	return d.Name
}

// Go 语言中，并不需要显式地声明实现了哪一个接口，只需要直接实现该接口对应的方法即可
func TestInterface(*testing.T) {
	// 实例化 Cat后，强制类型转换为接口类型 Animate
	var animal Animate
	animal = Cat{Name: "小猫"}
	fmt.Println(animal.getName())

	// 如果定义了一个没有任何方法的空接口，那么这个接口可以表示任意类型。
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m) // map[age:18 name:Tom scores:[98 99 85]]
}