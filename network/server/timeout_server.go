package server

import (
	"net"
	"strings"
	"time"
)

type TimeoutServer struct {
	is_working bool // lets add tumbler
}

func (s *TimeoutServer) TurnOn() {
	if !s.is_working {
		s.is_working = true
		go s.run()
	}
}

func (s *TimeoutServer) TurnOff() {
	s.is_working = false
}

func (s *TimeoutServer) run() {
	listner, err := net.Listen("tcp", "localhost:2525") // init server listner for localhost:4545 and tcp
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	for s.is_working {
		connection, err := listner.Accept() // lets accept any connection
		if err != nil {
			panic(err)
		}
		defer connection.Close()
		connection.SetReadDeadline(time.Now().Add(time.Second)) // lets add deadline: now + second

		readed := ByteRead(connection, 8)

		output := []byte(strings.ReplaceAll(string(readed), " ", "-"))
		ByteWrite(connection, output, 8)
	}
}

func ByteRead(connection net.Conn, batchSize int) []byte {
	var readed []byte
	for {
		buff := make([]byte, batchSize)
		n, err := connection.Read(buff)
		if err != nil {
			break
		}
		readed = append(readed, buff[:n]...)
		connection.SetReadDeadline(time.Now().Add(time.Millisecond * 200))
	}
	return readed
}

func ByteWrite(connection net.Conn, inputMessage []byte, batchSize int) {
	for i := 0; i < len(inputMessage); i += batchSize {
		if i+batchSize < len(inputMessage) {
			connection.Write(inputMessage[i:(i + batchSize)])
		} else {
			connection.Write(inputMessage[i:])
		}
	}
}
