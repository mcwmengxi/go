package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type timeHandle struct {
	mutex  sync.Mutex
	format string
}

func (timeHandle *timeHandle) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	timeHandle.mutex.Lock()
	defer timeHandle.mutex.Unlock()
	timeHandle.format = time.RFC1123
	currentTime := time.Now().Format(timeHandle.format)
	res.Write([]byte("current time: " + currentTime))
}

func main() {
	// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	// ServeMux 和 Handler
	mux := http.NewServeMux()
	// 调用HandlerFunc直接把一个函数注册到指定的路由上
	mux.HandleFunc("/test", handle)
	th := &timeHandle{format: time.RFC1123}
	// 调用Handle把一个Handler注册到路由里
	mux.Handle("/time", th)
	log.SetPrefix("[http学习] ")              // 给log内容添加前缀
	log.SetFlags(log.Ldate | log.Llongfile) // 设置日志的一系列标志

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", mux)) // Fatal执行完就退出程序了
}

func handle(res http.ResponseWriter, req *http.Request) {
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	res.Write([]byte("server response!"))
}
