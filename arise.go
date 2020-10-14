package main

import (
	"os"

	"github.com/arisetransfer/arise/client"
	"github.com/arisetransfer/arise/server"
)

func main() {
	var instr = os.Args[1]
	if instr == "relay" {
		server.StartRelay()
	} else {
		var codeOrFile = os.Args[2]
		if instr == "send" {
			client.Sender(codeOrFile)
		} else if instr == "recieve" {
			client.Reciever(codeOrFile)
		}
	}
}
