package main

import (
	"gorutineLearning/basics"
	"gorutineLearning/chanels"
	"gorutineLearning/waitGroup"
)

func main() {
	println("BASICS")
	basics.Basics()
	println("--------------------\nWaitgroup")
	waitGroup.WaitGroup()
	println("--------------------\nChanel")
	chanels.Chanels()
}
