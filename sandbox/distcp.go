package main

import "github.com/ShoresyLX/go-osc/osc"

func main() {
	addr := "127.0.0.1"
	port := "8001"

	t, _ := osc.NewTcpServer(addr, port)
	t.Dispatcher.AddMsgHandler("/eos/out/*", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})
	t.Connect()
	t.ServeForever()
}
