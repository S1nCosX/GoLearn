package server

import (
	"log"
	"net"
	"os"
)

type SimpleServer struct {
	is_working bool
}

func (s *SimpleServer) TurnOn() {
	if !s.is_working {
		s.is_working = true
		go s.run()
	}
}

func (s *SimpleServer) TurnOff() {
	s.is_working = false
}

func (s *SimpleServer) run() {
	listner, err := net.Listen("tcp", "localhost:4646") // init server listner for localhost:4545 and tcp
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

		readed := make([]byte, 128)
		n, err := connection.Read(readed)
		if err != nil {
			log.New(os.Stdout, "warning", 0)
			panic(err)
		}
		outputStr := []byte("you succesfly connected to server got your message: ")
		connection.Write(append(outputStr, readed[:n]...))
	}
}
