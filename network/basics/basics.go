package basics

import (
	"net"
	"networkLearn/server"
)

func Basics() {
	// firstly i think we must learn how to create simplest server
	// so in preperation i writen simplest server that got connection and after connection created write in some text
	var server server.SimpleServer
	server.TurnOn()
	defer server.TurnOff()

	connection, err := net.Dial("tcp", "localhost:4646")
	if err != nil {
		panic(err)
	}
	readed := make([]byte, 128) // connetion is kinda iostream, lets write something and got response
	connection.Write([]byte("biba boba flex"))
	n, err := connection.Read(readed)
	if err != nil {
		panic(err)
	}

	println(string(readed[:n]))
	connection.Close()
}
