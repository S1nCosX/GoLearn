package main

import (
	"networkLearn/basics"
	httpLearn "networkLearn/http"
	"networkLearn/timeout"
)

func main() {
	basics.Basics()
	timeout.Timeout()
	httpLearn.Http()
}
