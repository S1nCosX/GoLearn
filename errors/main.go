package main

import (
	"errors"
	"fmt"
)

type BibaError struct{}

func (this BibaError) Error() string {
	return "Oh ooooh u got trouble yang man"
}

func testFunc(i int) (int, error) {
	if i != 0 {
		return 1, BibaError{}
	}
	return 0, nil
}
func testFuncForErrors() (int, error) {
	return 1, errors.New("this is error created with errors lib")
}

func recoverFunction() {
	for i := 0; i < 5; i++ {
		fmt.Println(recover(), i)
	}
}

func main() {
	// error actualy is more interface than type.
	// I write my own minimal error, to throw it. Take a look higher.
	// lets try it.
	res, err := testFunc(1)
	fmt.Println(res, err)
	// also actualy nil can be writen in error
	res, err = testFunc(0)
	fmt.Println(res, err)
	// also there lib named errors. That used for fast
	res, err = testFuncForErrors()
	fmt.Println(res, err)

	// in this theme nice to know about some code words like defeк, panic and recover
	// defer is a word that as function into view range close stack.
	// that means, if we exit from {} construction, we'll execute functions in LIFO order. Let me demonstrate
	defer fmt.Println("this defer added first")
	defer fmt.Println("this defer added second")
	defer fmt.Println("this defer added third")
	// lets add recover. Recover is a standart func like len or cap, that let us get value of error thrown by panic. This can be used to create try-catch logic in code
	defer recoverFunction() // as you can see, recover can provide us error only 1 time, all next calls of recover will return nil
	panic(errors.New("FUNNY ERR HAHA"))
}
