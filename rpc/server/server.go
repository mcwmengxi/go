package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	c := new(Calc)
	rpc.Register(c)
	rpc.HandleHTTP()
	listen, e := net.Listen("tcp", ":8088")
	if e != nil {
		log.Fatal(e.Error())
	}
	http.Serve(listen, nil)
}

type Req struct {
	Num1 int
	Num2 int
}
type Res struct {
	Total int
}
type Calc int

func (c *Calc) SumTotal(res Res, req *Req) error {
	res.Total = req.Num1 + req.Num2
	return nil
}
