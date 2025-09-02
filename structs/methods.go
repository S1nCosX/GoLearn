package main

import "fmt"

type SomeStruct struct {
	array [10]int
	sum int
}

func (ss SomeStruct) someMeth() (output int) { //there we'll create copy of an struct, not  using actual struct obj
	output = 0
	for _, val := range(ss.array) {
		output += val
	}
	ss.sum = output
	return output
}

func (ss *SomeStruct) someMethWithPointer() (output int) { //there we'll use pointer on struct
	output = 0
	for _, val := range(ss.array) {
		output += val
	}
	ss.sum = output
	return output
}

func we_will_talk_about_METHODS() {
	fmt.Println("--------------\nTODAY WE'LL TALKIN BOUT METHODS--------------\n")

	structExemp := SomeStruct{}

	for i := 0; i < 10; i++ {
		structExemp.array[i] = 2 * i
	}
	fmt.Println("method will sum aray but not change it's sum val couse used copy of object", structExemp.someMeth(), "struct:", structExemp)
	fmt.Println("method will sum aray", structExemp.someMethWithPointer(), "struct:", structExemp)
}