// 使用 ProviderSet 进行分组
package prov

import "github.com/google/wire"

// var ProvSet = wire.NewSet(NewMessage, NewGreeter, NewEvent)
// 使用 Struct 定制 Provider
// var ProvSet = wire.NewSet(wire.Struct(new(Message), "*"), NewGreeter, NewEvent)

// 使用 Struct 定制 Provider
var ProvSet = wire.NewSet(NewMessage, wire.FieldsOf(new(*Message), "Msg")) 
