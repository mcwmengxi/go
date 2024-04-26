package channeldemo

import (
	"fmt"
	"time"
)

func main()  {
	//chan_1()
	fmt.Println("--------")
	//chan_2()
	fmt.Println("--------")
	chan_3()
	fmt.Println("--------")
	chan_4()
	fmt.Println("--------")
	// chan_5()
}

// chan通道 类似队列，先进先出
func chan_1()  {
	fmt.Println("主线程开始：", )
	go func() {
		fmt.Println("chan通道线程")
	}()
	// chan通道是并发的线程机制，而不是并行的线程机制
	// 通过主线程休眠让其他线程执行
	time.Sleep(1 * time.Second)
	fmt.Println("主线程结束：", )
}
func chan_2()  {
	// 不带缓存的通道
	// ch_1 := make(chan string)
	// 无缓冲区的通道直接写入会造成死锁
	// ch_1 <- "data1"
	// 带10个缓存的通道
	ch_2 := make(chan string, 5)
	ch_2 <- "a" // 入 chan
	ch_2 <- "a"
	ch_2 <- "a"
	time.Sleep(1 * time.Second)
	go func() {
		// 从通道中取
		// val1 := <- ch_1
		// fmt.Println(ch_1, val1)
		// close(ch_1)
		val := <- ch_2 // 出 chan
		fmt.Println(val)
	}()



}
func chan_3()  {
	fmt.Println("main start")
	// 声明只读通道
	// ch_3 := make(<-chan string, 10)
	// val_3 := <- ch_3 // 写入报错all goroutines are asleep - deadlock!
	go func() {
		// fmt.Println("只读通道: ", val_3)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
func chan_4()  {
	fmt.Println("main start")
	// 声明只写通道
	ch_4 := make(chan<- string, 10)
	ch_4 <- "a"
	ch_4 <- "a"
	close(ch_4)
	// ch_4 <- "a"// 关闭通道之后再写会报panic: send on closed channel
	go func() {
		// invalid operation: cannot receive from send-only channel ch_4
		// 不能读取只写通道
		// val_4 := <- ch_4 
		// fmt.Println("只写通道: ", val_4)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func chan_5()  {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)
	go customer(ch) // 如果不提供消费者，通道就会阻塞

	time.Sleep(1 * time.Second)
	fmt.Println("main end")	
}

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func customer(ch chan string) {
	for {
		msg := <- ch
		fmt.Println(msg)
	}
}
