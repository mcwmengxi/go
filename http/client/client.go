package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// 用于控制代理、TLS 配置、保持活动状态、 压缩和其他设置，创建传输
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	res1, _ := client.Get("http://localhost:8080/test")
	res, _ := client.Get("http://localhost:8080/time")
	// 客户端必须在完成响应正文后关闭响应正文
	defer res.Body.Close()
	defer res1.Body.Close()
	headers := res.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	fmt.Printf("res status %s,statusCode %d\n", res.Status, res.StatusCode)

	fmt.Printf("res Proto %s\n", res.Proto)

	fmt.Printf("res content length %d\n", res.ContentLength)

	fmt.Printf("res transfer encoding %v\n", res.TransferEncoding)

	fmt.Printf("res Uncompressed %t\n", res.Uncompressed)
	b, _ := io.ReadAll(res.Body)
	fmt.Println("res", res.Status, string(b))
	fmt.Println("res1 header", res1.Header)
}
