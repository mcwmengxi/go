package main

import (
	"log"
	"net/http"
	"time"
)

type timeHandle struct {
	format string
}

func (th *timeHandle) ServeHttp(res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now().Format(th.format)
	res.Write([]byte("current time: " + currentTime))
}

func main() {
	// ServeMux 和 Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handle)
	th := &timeHandle{format: time.RFC1123}
	mux.Handle("/time", th)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handle(res http.ResponseWriter, req *http.Request) {
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	res.Write([]byte("server response!"))
}
