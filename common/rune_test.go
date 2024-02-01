package common

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	s1 := "小凶许"
	// 字符串的修改要转成rune切片
	s2 := []rune(s1)
	s2[0] = '大'
	fmt.Println(string(s2), s1)
}
