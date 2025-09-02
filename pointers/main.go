package main

import "fmt"

func tryingToChangeValueWithNoPointer(x int) {
	x = x * x
}

func tryingToChangeValueWithPointer(x *int) {
	*x = (*x) * (*x)
}

func main() {
	var v int = 10 //it is variable
	var p *int //it is pointer
	p = &v //& is key symbol to get pointer. There we set pointer equals pointer at V
	fmt.Println("variable:", v, "| pointer:", p)
	//if you want to work with pointer value, use *. 
	fmt.Println("pointers value:", *p)
	//also it is works for changing value at pointer
	//this is right way to work with variable under pointer
	*p += 1
	fmt.Println("pointers value:", *p, " | variable:", v)
	//this is wrong way, cause you changing address at memory, not value
	// p += 1
	// fmt.Println("pointers value:", *p, " | variable:", v)
	// code higher throws exception

	//there how you can work with funcs
	x := 2
	tryingToChangeValueWithNoPointer(x)
	fmt.Println("x without using pointer:", x)
	tryingToChangeValueWithPointer(&x)
	fmt.Println("x with pointer:", x)
}