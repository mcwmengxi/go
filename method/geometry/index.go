package geometry

import (
	"math"
)

type Point struct{ X, Y float64 }

// 包级别的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 选择器的方法和字段一般不能相同
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path代表一个线段的集合
// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (p Path) Distance() float64 {
	var sum float64
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func PathDistance(p Path) float64 {
	return p.Distance()
}
