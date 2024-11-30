//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
// 构建标记 条件编译 go run . 执行的是 main.go 和 wire_gen.go 两个文件，会忽略 wire.go
package main

import (
	prov "quickstart/wire/provides"

	"github.com/google/wire"
)

// wireApp init application.
// func wireApp(msg string, code int) (prov.Event, error) {
// 	// wire.Build(prov.NewMessage, prov.NewGreeter, prov.NewEvent)
// 	// wire.Build(prov.ProvSet)
// 	wire.Build(prov.NewMessage, wire.FieldsOf(new(*prov.Message), "Msg"))
// 	return prov.Event{}, nil
// }
func wireApp() prov.Message {
	// 假设没有提供 NewMessage，可以直接绑定值并返回
	wire.Build(wire.Value(prov.Message{
		Msg: "Binding Values",
		Code:    1,
	}))
	return Message{}
}

