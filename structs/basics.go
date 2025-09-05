package main

import "fmt"

// struct is just pseudoname for anonymus struct
type commonStruct struct {
	intField int
	stringField string
	arrayField [3]int
}

// structs can have anonymus fields, wierd flex but OK
type structWithAnonymusFields struct {
	intField int // you can have only 1 anonymus field at once in struct
	int
	string
}

func we_will_talk_about_BASICS() {
	fmt.Println("--------------\nTODAY WE'LL TALKIN BOUT BASICS OF USING STRUCTS\n")
	//struct init
	commonStructVariable := commonStruct{intField : 1, stringField : "aboba", arrayField : [3]int{1, 2, 3}}
	fmt.Println(commonStructVariable)

	fmt.Println("int field:", commonStructVariable.intField, "| string field:", commonStructVariable.stringField, "| array field:", commonStructVariable.arrayField)

	//there such a thing as anonymus struct idk what purpus it is serves but for protocol
	anonymusStruct := struct {
		Name string
		Age int
		BurnDate [3]int
	} {"Adolf", 43, [3]int{11, 4, 1870}}
	fmt.Println(anonymusStruct)
	// you can use anonymus structs for slices and arrays as well
	sliceOfAnonymStruct := []struct {
		int
		string
	} {{1, "2"}, {3, "4"}, {5, "6"}}
	fmt.Println(sliceOfAnonymStruct)

	// init struct with anonymus fields
	variableWithAnonymusFields := structWithAnonymusFields{12, 1, "ABOBUS"}
	fmt.Println(variableWithAnonymusFields)
	fmt.Println("int field:", variableWithAnonymusFields.intField, "| anonymus int field:", variableWithAnonymusFields.int, "| anonymus string field", variableWithAnonymusFields.string)

	// copy
	someStructCopy := variableWithAnonymusFields
	someStructCopy.intField = 1
	someStructCopy.string = "boba"
	fmt.Println(variableWithAnonymusFields, someStructCopy)
}