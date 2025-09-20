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
	// Close method talks itself about its destination. We can close chanel and after this nobody will have ability to write in it
	intBufferedChanel := make(chan int, 3)
	intBufferedChanel <- 1
	intBufferedChanel <- 2
	close(intBufferedChanel)
	// If we'll try to put some value after closing chanel, it'll cause exeption
	// IntBufferedChanel <- 3
	// So what we can do now?
	// We can get info about closed our chanel or not, when getting data from it
	data, isOpened := <-intBufferedChanel
	println("so chanel info:", data, isOpened)
	// This used to message consumers: There no new data here, get last messages and getout.
	wg.Add(1)
	go func(chanel chan int, wg *sync.WaitGroup) {
		for val := range chanel { // This loop wi'll works until chanel wi'll be closed
			println("readed some data from chanel:", val)
		}
		wg.Done()
	}(intBufferedChanel, &wg)
	wg.Wait()
	// also we can create read|write only chanels:
	boolChan := make(chan bool)
	go func(boolChan chan<- bool) {
		boolChan <- true
		// There we cannot get val from chanel. Line like in previous lambda will cause problems
	}(boolChan)
	go func(boolChan <-chan bool) {
		// So there we cannot write boolChanel <- true. It'll cause exeption
		println("test function with read only chanel got:", <-boolChan)
	}(boolChan)
	// Select... this was kinda tricky. this is switch-case like construction where we can create some breanching for expection of data from chanels
	// Select is mechanism similar to getting data from chanel like: <- Chanel, it'll stop thread until one of expected chanels send a signal about being ready
	// Also there option that give us ability to not stop thread for waiting ready signal: default. if there default in select construction instead of waiting we'll
	// continue with execution code in default section
	// example:
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func(strChan chan<- int) {
		time.Sleep(time.Second * 2)
		strChan <- 21
	}(chan1)
	go func(boolChan chan<- int) {
		time.Sleep(time.Second * 2)
		boolChan <- 12
	}(chan2)

	println("Lets wait until select worked")
	select {
	case inp := <-chan1:
		println("got message from bool chan:", inp)
	case inp := <-chan2:
		println("got message from str chan:", inp)
	}
	// as you can see in this case we were waiting until any data thrown in chanels
	// now with default
	chan11 := make(chan int)
	chan22 := make(chan int)
	go func(strChan chan<- int) {
		time.Sleep(time.Second * 2)
		strChan <- 21
	}(chan11)
	go func(boolChan chan<- int) {
		time.Sleep(time.Second * 2)
		boolChan <- 12
	}(chan22)

	println("Lets wait until select worked, now added default")
	select {
	case inp := <-chan11:
		println("got message from bool chan:", inp)
	case inp := <-chan22:
		println("got message from str chan:", inp)
	default:
		println("there are no messages")
	}
	// also there thing, it has no hierarhy. If you need one, do select in default section of select, like:
	// select {
	//		case chan1:
	//		default: select {
	//			case chan2:
	//		}
	// }
	// also closed chanel constantly message ready
	chan111 := make(chan int)
	close(chan111)
	select {
	case <-chan111:
		println("as you see chan111 is closed, have no data, but we triggered case chan111")
	}
}
