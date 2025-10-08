package timeout

import (
	"net"
	network "networkLearn/server"
	"time"
)

func Timeout() {
	// so what if we want to sand multiple packages throu connection?
	// we must setup timeouts if we have system require fast as request/response or skip
	var server network.TimeoutServer
	server.TurnOn()
	defer server.TurnOff()

	connection, err := net.Dial("tcp", "localhost:2525")

	if err != nil {
		panic(err)
	}
	defer connection.Close()
	connection.SetDeadline(time.Now().Add(time.Second))

	network.ByteWrite(connection, []byte("spaces in this message converted into dash"), 8)

	readed := network.ByteRead(connection, 8)
	println(string(readed))
}
