package channeldemo

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(1 * time.Second)
	ch <- url // 将 url 发送到信道	
}

func TestChannel(t *testing.T) {
	for i := 0; i < 3; i++ {
		go download("https://www.baidu.com/%s" + strconv.Itoa(i))
		// fmt.Println(<- ch) // 从信道中接收
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息。
		fmt.Println("finish", msg)
	}
}
