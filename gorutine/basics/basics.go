package basics

import "time"

func Sum(iteration, i int, sum *int) {
	for j := 0; j <= i; j++ {
		*sum += j
		println("run num", iteration, "core", i, "added", j, "result sum now is", *sum)
	}
}

func Basics() {
	// Gorutines are realy easy to start. Just write go before func and it'll start in another thread
	sum := 0
	expected := 0
	for i := 0; i < 5; i++ { // i was too lazy to calculate myself so this what we expect, but faster cause multithreading
		for j := 0; j <= i; j++ {
			expected += j
		}
	}
	for i := 0; i < 5; i++ {
		go Sum(1, i, &sum)
	}
	println("First part where we dont wait for threads finished result sum we achived", sum, "/", expected)
	// we have unexpected behavior because of runtime of new threads and runtime of main thread wich is thread where we starts main.
	// main -> Basics. if Basics ended earlier then sum threads we'll get random results.
	// so lets wait a bit for result:
	niceSum := 0
	for i := 0; i < 5; i++ {
		go Sum(2, i, &niceSum)
	}
	time.Sleep(time.Millisecond * 200)
	println("First part where we waited 1s for threads finished result sum we achived", niceSum, "/", expected)
	// also gorutine not ask you to make your func named. Lambda is good as well
	for i := 0; i < 5; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				println("this is message from lamda", i, "it iterate j", j)
			}
		}(i)
	}
	time.Sleep(time.Millisecond * 200)
}
