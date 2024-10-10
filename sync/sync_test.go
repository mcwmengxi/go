package sync

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

var (
	counter int
	mutex   sync.Mutex
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


var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(1 * time.Second)
	wg.Done()
}

func TestWaitGroup(t *testing.T) {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("https://www.baidu.com/%s" + strconv.Itoa(i))
	}
	wg.Wait()
	fmt.Println("Done!")
}


func TestMutex(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementCounter()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

func incrementCounter() {
	mutex.Lock()
	defer mutex.Unlock()

	counter++
}