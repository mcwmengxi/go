package now

import (
	"fmt"
	"testing"
)

var n = Now{}

func TestNow(t *testing.T) {
	fmt.Println(n.String())
}

func TestBeginningOfMinute(t *testing.T) {
	fmt.Println(n.beginningOfMinute().String())
}
func TestBeginningOfHour(t *testing.T) {
	fmt.Println(n.BeginningOfHour().String())
}

func TestBeginningOfDay(t *testing.T) {
	fmt.Println(n.BeginningOfDay().String())
}
func TestNowBeginningOfWeekday(t *testing.T) {
	fmt.Println(n.BeginningOfWeek())
}
func TestBeginningOfMonth(t *testing.T) {
	fmt.Println(n.BeginningOfMonth().String())
}
