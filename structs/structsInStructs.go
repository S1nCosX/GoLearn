package main

import "fmt"

type Person struct {
	name string
	age int
}

type DataWithNestedPerson struct {
	login string
	password string
	nestedPerson Person
}

type DataWithEmbeddedPerson struct {
	login string
	password string
	Person
}

func we_will_talk_about_STRUCTS_WITHIN_STRUCTS() {
	fmt.Println("--------------\nTODAY WE'LL TALKIN BOUT STRUCTS WITHIN OTHER STRUCTS--------------\n")
	personExemp := Person{name: "Biba", age: 100}
	fmt.Println("we have person:", personExemp)
	
	nestedPersonExemp := DataWithNestedPerson{login: "amogus", password: "password", nestedPerson: personExemp}
	fmt.Println("we can nest him in some struct. if struct field have name it is nested struct", nestedPersonExemp)

	embeddedPersonExamp := DataWithEmbeddedPerson{"abobus", "password2", personExemp}
	fmt.Println("also if struct within struct is noname it is embedded struct", embeddedPersonExamp)
}

type Node struct{
	value int
	next *Node
} // kinda list node

func we_will_talk_about_POINTERS_ON_STRUCTS_WITHIN_STRUCTS() {
	fmt.Println("--------------\nTODAY WE'LL TALKIN BOUT POINTERS ON STRUCTS WITHIN OTHER STRUCTS--------------\n")
	first := Node{value: 1}
	second := Node{value: 2}
	first.next = &second
	third := Node{value: 3}
	second.next = &third
	//list created
	curr := &first
	for ;curr != nil; { // si we'll iterate while we are not in nil, that means we reached end of list
		fmt.Println(curr.value)
		curr = curr.next
	}
}