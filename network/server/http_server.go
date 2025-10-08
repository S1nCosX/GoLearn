package server

import (
	"context"
	"net/http"
)

type HttpServer struct {
	is_working bool
}

func (s *HttpServer) TurnOn() {
	if !s.is_working {
		s.is_working = true
		go s.run()
	}
}

func (s *HttpServer) TurnOff() {
	s.is_working = false
}

func (s *HttpServer) run() {
	// don't use http lib like this because this line will put execution in infinite cycle of listening port until got any error
	// err := http.ListenAndServe(":6060", nil)
	// but how we will do this.
	// firstly lets create server:
	httpServer := http.Server{
		Addr: ":3737", // addres for handling
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				return
			}
			n, err := w.Write([]byte("Got your get request"))
			if err != nil {
				print("tryed to write", n, "bytes")
				panic(err)
			}
		}),
	}

	go func(s *HttpServer, actualServer *http.Server) {
		for s.is_working {
		}
		if s != nil {
			actualServer.Shutdown(context.Background())
		}
	}(s, &httpServer)

	err := httpServer.ListenAndServe()

	print("succesfuly shutted down server")
	if err != http.ErrServerClosed {
		s.is_working = false
		panic(err)
	}
}
