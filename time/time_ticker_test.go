package array

import (
	"fmt"
	"testing"
	"time"
)

// 循环触发
func TestTicker(t *testing.T) {
	timeTicker := time.NewTicker(1 * time.Second)
	count := 0

	for {
		select {
		case tim := <-timeTicker.C:
			fmt.Println(tim.Format(time.RFC3339))
			count++
			if count == 10 {
				fmt.Println("ticker stop")
				timeTicker.Stop()
				return
			}
		}
	}
}
