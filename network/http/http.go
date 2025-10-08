package httpLearn

import (
	"io"
	"net/http"
	"networkLearn/server"
	"time"
)

func Http() {
	server := server.HttpServer{}
	server.TurnOn()
	defer server.TurnOff()
	time.Sleep(time.Second)
	response, err := http.Get("http://:3737")
	if err != nil {
		panic(err)
	}
	buff := make([]byte, 1024)
	for {
		n, err := response.Body.Read(buff)
		if n == 0 {
			break
		}
		println(string(buff[:n]))
		if err != nil {
			if err == io.EOF {
				response.Body.Close()
				break
			}
			panic(err)
		}
	}
}
