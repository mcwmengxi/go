package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

//	type Person struct {
//		Name   string
//		Age    int64
//		Weight float64
//	}
type Person struct {
	ID       int64   `json:"id,string"`
	Score    float64 `json:"score,string"` // 添加string tag
	Name     string  `json:"name"`         // 指定json序列化/反序列化时使用小写name
	Password string  `json:"password"`
	Age      int64
	Weight   float64                    `json:"-"` // 指定json序列化/反序列化时忽略此字段
	Hobby    []string                   `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"` // 嵌套的结构体为空值时，忽略该字段,使用结构体指针+omitempty tag
}
type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

type PublicPerson struct {
	*Person
	Password *struct{} `json:"password,omitempty"` // 不修改原结构体忽略空值字段
}

func TestBasic(t *testing.T) {
	p1 := Person{
		Name:     "小明",
		Age:      18,
		Weight:   71.5,
		Password: "123456",
	}
	// struct -> json string
	// b, err := json.Marshal(p1)
	b, err := json.Marshal(PublicPerson{Person: &p1})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

	// json string -> struct
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)

	jsonStr1 := `{"id": "1234567","score": "88.50"}`
	var c1 Person
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}

}
