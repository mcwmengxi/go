package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncDemo(t *testing.T) {
	OnceD()
	OnceD()
}

var once sync.Once

func OnceD() {
	once.Do(func() {
		fmt.Println("OnceD123")
	})
}
