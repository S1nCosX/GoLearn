package mutex

import (
	"sync"
)

func Mutex() {
	// mutex is sync mechanism that block shared variables.
	// actualy this kinda stupid. I know that mutex is important func, but I cant imagine better demonstration than Fibonacci num calculation
	var mu sync.Mutex
	var wg sync.WaitGroup
	curr := 1
	prev := 1
	wg.Add(2)
	fib := func(num int, curr, prev *int, wg *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			next := *curr + *prev
			*prev = *curr
			*curr = next
			println("no mu", num, ":", *curr)
		}
		wg.Done()
	}
	go fib(1, &curr, &prev, &wg)
	go fib(2, &curr, &prev, &wg)
	wg.Wait()
	//okay... actualy this code must cause troubles... but i couldnt catch this...\
	curr = 1
	prev = 1
	mufib := func(num int, curr, prev *int, wg *sync.WaitGroup, mu *sync.Mutex) {
		for i := 0; i < 5; i++ {
			mu.Lock()
			next := *curr + *prev
			*prev = *curr
			*curr = next
			mu.Unlock()
			println("with mu", num, ":", *curr)
		}
		wg.Done()
	}
	wg.Add(2)
	go mufib(1, &curr, &prev, &wg, &mu)
	go mufib(2, &curr, &prev, &wg, &mu)
	wg.Wait()
}
