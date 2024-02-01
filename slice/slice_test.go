package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// slice
	s1 := make([]int, 5, 10)
	fmt.Printf("s1：%v len(s1)：%d cap(s1)：%d\n", s1, len(s1), cap(s1))
	s1 = []int{1, 2, 3}
	s2 := s1 //s1 s2都指向同一个底层数组
	s2[0] = 2
	fmt.Println(s2, s1)

	s3 := make([]any, 2, 2)
	fmt.Printf("s3：%v len(s3)：%d cap(s3)：%d\n", s3, len(s3), cap(s3))
	s3 = append(s3, "小凶许")
	fmt.Printf("s3：%v len(s3)：%d cap(s3)：%d\n", s3, len(s3), cap(s3))
	s4 := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		s4 = append(s4, i)
	}
	fmt.Println(s4)
}

func TestArray(t *testing.T) {
	// array
	s1 := [3]int{1, 2, 3}
	s2 := s1
	s2[0] = 2
	fmt.Println(s2, s1)
}
