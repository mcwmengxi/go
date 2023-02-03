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

func (timeHandle timeHandle) ServeHttp(res http.ResponseWriter, req *http.Request) {
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
	// mux.HandleFunc("/test", handle)
	// th := &timeHandle{format: time.RFC1123}
	mux.Handle("/time", timeHandle{})
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handle(res http.ResponseWriter, req *http.Request) {
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	res.Write([]byte("server response!"))
}
