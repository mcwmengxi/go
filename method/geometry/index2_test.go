package geometry

import (
	"fmt"
	"testing"
)

// 基于指针对象的方法
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func TestScaleBy(t *testing.T) {
	p := Point{1, 2}
	p.ScaleBy(2)
	fmt.Println(p)
	if p.X != 2 || p.Y != 4 {
		t.Errorf("ScaleBy failed, got %v", p)
	}
}
