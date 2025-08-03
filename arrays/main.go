package main

import "fmt"

func main() {
	//init array
	var intArray [10]int = [10]int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 } //you can init array with some kind of value
	fmt.Println("array: ", intArray)

	var boolArray = [...]bool{ true, false, true} //you can give implicit size with [...]
	fmt.Println("array where unknown size: ", boolArray)

	floarArray := [...]int{ 1, 2, 3, 4, 5 } //also there are implicit creation 
	fmt.Println("implicit array:", floarArray)

	var stringMultiArray [5][5]string = [5][5]string{
		{"ab", "bc", "cd", "de", "ef"},
		{"Ab", "Bc", "Cd", "De", "Ef"},
		{"aB", "bC", "cD", "dE", "eF"},
		{"AB", "BC", "CD", "DE", "EF"},
		{"biba", "boba", "|", "dva", "dolboyeba"},
	}
	fmt.Println("multi array:", stringMultiArray) 

	var intMultiArray = [3][2]int{
		{1, 2},
		{2},
	} // if you keep some values uninitiated, it'll have default value
	fmt.Println("array with uninited vals:", intMultiArray) 
	
	//length of array
	var arrayForLength = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("\nArray: ", arrayForLength, "\nhave length:", len(arrayForLength))

	//copy
	var arrayToCopy = [...]int{1, 2}
	copiedArray := arrayToCopy
	fmt.Println("\narray to copy:", arrayToCopy, "\ncopied array:", copiedArray)
	
	copiedArray[0] = 11 // it is a copy not a link
	fmt.Println("\narray to copy:", arrayToCopy, "\ncopied array:", copiedArray)
}