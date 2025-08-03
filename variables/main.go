package main

import "fmt"

func main () {
	//var is for Variable can change it value during runtime, but cannot change type

	var intVar int64 //have variants as 8 16 32 64, also uint for unsigned. 
	intVar = 10
	fmt.Println("int variable:", intVar)

	intVar = 12
	fmt.Println("changed int variable:", intVar)

	var floatVar float64 //have 32 64 variants (float/double in other lang)
	floatVar = 12.
	fmt.Println("float variable:", floatVar)

	var complexVar complex64 // have 64 128 variants where half is for imaginary and half for real (idk what purpuse i'll need that...)
	complexVar = 15 + 10i
	fmt.Println("complex variable:", complexVar)

	var boolVar bool //bool is bool. no variants
	boolVar = false
	fmt.Println("boolean variable:", boolVar)

	var stringVar string //string is string there are no such a thing as char.
	stringVar = "'this is string'"
	fmt.Println("string variable:", stringVar)

	//also there are implicit creation of variable (i think it's kinda unreadable btw) and btw you cannot change its type in runtime
	var implicitVar = 10 //can be init with var
	fmt.Println("created with var implicit variable:", implicitVar)

	equalsImplicitVar := 20 //can be init with :=
	fmt.Println("created with := implicit variable:", equalsImplicitVar)

	//constants are constans
	const constant int = 21
	fmt.Println("int constant:", constant)

	const implicitConstant = 12123 + 21398i //can be implicit
	fmt.Println("implicit constat:", implicitConstant)
}