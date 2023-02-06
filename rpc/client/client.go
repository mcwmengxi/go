package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}
type Res struct {
	Total int
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:8088")
	req := &Req{9, 10}
	// args := &server.Args{7,8}
	// var sum int
	// client.Call("Calc.SumTotal", res, &sum)
	res := new(Res)
	c := client.Go("Calc.SumTotal", req, &res, nil)
	val := <-c.Done
	fmt.Println("total", req, res, val)
}
