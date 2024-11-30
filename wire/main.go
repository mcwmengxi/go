package main

// func InitializeEvent() prov.Event {
// 	message := prov.NewMessage()
// 	greeter := prov.NewGreeter(message)
// 	event := prov.NewEvent(greeter)
// 	return event
// }


func main() {
	// event := InitializeEvent()
	event, _ := wireApp("dfas", 23)
	event.Start()
	// SuperSet = wire.NewSet(foobarbaz.ProvideFoo, foobarbaz.ProvideBar, foobarbaz.ProvideBaz)

}
