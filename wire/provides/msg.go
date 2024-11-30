package prov

// type Message string

// // 接收参数作为消息内容
// func NewMessage(msg string) Message {
// 	return Message(msg)
// }
type Msg string

type Message struct {
	Code int
	Msg Msg
}
// NewMessage 注意，这里返回的是指针类型
func NewMessage(content string, code int) *Message {
	return &Message{
		Msg: Msg(content),
		Code:    code,
	}
}
