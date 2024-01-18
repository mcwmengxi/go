package array

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeAfter(t *testing.T) {
	fmt.Println("before: " + getNow())
	select {
	case now := <-time.After(3 * time.Second):
		fmt.Println("delay 3 seconds")
		fmt.Println("after: " + time.Unix(now.Unix(), 0).Format(time.RFC3339))
	}
}

func getNow() string {
	return time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)
}
