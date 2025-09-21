package main

import (
	"gorutineLearning/basics"
	"gorutineLearning/chanels"
	"gorutineLearning/mutex"
	"gorutineLearning/waitGroup"
)

func main() {
	println("BASICS")
	basics.Basics()
	println("--------------------\nWaitgroup")
	waitGroup.WaitGroup()
	println("--------------------\nChanel")
	chanels.Chanels()
	println("--------------------\nMutex")
	mutex.Mutex()
}
