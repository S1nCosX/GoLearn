package chanels

import (
	"sync"
	"time"
)

func Chanels() {
	// chanels in go is kind of type that empliments logic of wait and pointer on value in one
	// code word: chan (chan <TYPE>)
	// you can init it with make(chan <TYPE>) no other options if chanel is not initialisated it cant be used
	strChan := make(chan string)
	// syntaxis for put and get from chanel is realy simple <- chanel_name <-
	// wait is on both sides as you may see on this example
	go func(stringChanel chan string) {
		time.Sleep(time.Second)
		strChan <- "amogus"
		strChan <- "aboltus"
		println("lamda end")
	}(strChan)
	println(<-strChan)
	time.Sleep(time.Second)
	println(<-strChan)
	println("chanel blocking check ended")
	// also we can create buffered chanel, that means that we can send more than one messache through chanel
	// this is nice if our sender can create more messages than can be processed with one thread
	dataChan := make(chan struct {
		int
		string
	}, 20)
	var wg sync.WaitGroup
	wg.Add(3)
	consumer := func(consumerNum int, dataChan chan struct {
		int
		string
	}, wg *sync.WaitGroup) {
		for data := range dataChan {
			encodedStr := []rune{}
			for _, v := range []rune(data.string) {
				encodedStr = append(encodedStr, rune(data.int)+v)
			}
			println("Data encoder", consumerNum, "got message", data.string, "must be coded with offset", data.int, "and coded it into", string(encodedStr))
		}
		wg.Done()
	}
	for i := 0; i < 20; i++ {
		dataChan <- struct {
			int
			string
		}{int: i, string: "Alabuga polileh hehe"}
	}
	go consumer(1, dataChan, &wg)
	go consumer(2, dataChan, &wg)
	go consumer(3, dataChan, &wg)
	// chanels in golang are thread safe, but we'll get deadlock without 1 line.
	// Because our consumers reads until get signal from chanel that threre will be no new data.
	// so here we'll learn about close
	close(dataChan)
	wg.Wait()
	// close method talks itself about its destination. We can close chanel and after this nobody will have ability to write in it
	intBufferedChanel := make(chan int, 3)
	intBufferedChanel <- 1
	intBufferedChanel <- 2
	close(intBufferedChanel)
	// if we'll try to put some value after closing chanel, it'll cause exeption
	// intBufferedChanel <- 3
	// so what we can do now?
	// we can get info about closed our chanel or not, when getting data from it
	data, isOpened := <-intBufferedChanel
	println("so chanel info:", data, isOpened)
	// this used to message consumers: There no new data here, get last messages and getout.
	wg.Add(1)
	go func(chanel chan int, wg *sync.WaitGroup) {
		for val := range chanel { // this loop wi'll works until chanel wi'll be closed
			println("readed some data from chanel:", val)
		}
		wg.Done()
	}(intBufferedChanel, &wg)
	wg.Wait()
	// also we can create read|write only chanels:
	boolChan := make(chan bool)
	go func(boolChan <-chan bool) {
		// so there we cannot write boolChanel <- true. It'll cause exeption
		println("test function with read only chanel got:", <-boolChan)
	}(boolChan)
	go func(boolChan chan<- bool) {
		boolChan <- true
		// there we cannot get val from chanel. Line like in previous lambda will cause problems
	}(boolChan)
}
