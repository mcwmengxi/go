package main

import (
	"log"
	"net/http"
)

func main() {
	// ServeMux 和 Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handle)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handle(res http.ResponseWriter, req *http.Request) {
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	res.Write([]byte("server response!"))
}
