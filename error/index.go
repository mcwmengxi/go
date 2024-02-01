package main

import (
	"errors"
	"fmt"
)

var Crashed = errors.New("crashed")

func main() {
	err := Crashed
	if err != nil {
		fmt.Println(err)
	}
}
