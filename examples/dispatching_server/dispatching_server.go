package main

import "github.com/ShoresyLX/go-osc/osc"

func main() {
	addr := "127.0.0.1:9002"

	d := osc.NewStandardDispatcher()
	d.AddMsgHandler("/eos/out/*", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})
	server := &osc.Server{
		Addr:       addr,
		Dispatcher: d,
	}

	server.ListenAndServe()
}
