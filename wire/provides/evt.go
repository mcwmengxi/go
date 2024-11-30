package prov

import (
	"errors"
	"fmt"
	"time"
)

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) (Event, error) {
		// 模拟创建 Event 报错
		if time.Now().Unix()%2 == 0 {
			return Event{}, errors.New("new event error")
		}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg.Msg)
}
