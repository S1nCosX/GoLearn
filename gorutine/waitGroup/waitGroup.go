package waitGroup

import (
	"gorutineLearning/basics"
	"sync"
)

func WaitGroup() {
	// this is part where we have question. Is there kinda threadpool or something like this, where we can manage is threads done?
	// and answer is "ofcourse, dummie" there all sync lib for sinc our threads, but this part is only for waitgroup
	// waitgroup is object that hold in it variable for ammount of threads we starts
	// with func Add(num) we can add amount of threads we are waiting for
	// so lets make our pervious task but without sticks and weels
	var wg sync.WaitGroup
	wg.Add(5)
	sum := 0
	println("lets make it right, i set this part under waitgroup Wait")
	for i := 0; i < 5; i++ {
		go func(i int, sum *int, wg *sync.WaitGroup) {
			basics.Sum(0, i, sum)
			wg.Done()
		}(i, &sum, &wg)
	}
	// So if we want to make our thread stop until job in other threads done, we use waitgroup.Wait
	wg.Wait()
	println("Now job is done and we got", sum, "/", 20)
	// restart for program is fast, so you can check stability of this construction.
}
