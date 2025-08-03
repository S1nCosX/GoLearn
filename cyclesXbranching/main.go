package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { //for is just like in c++
		fmt.Println("in for i =", i)
	}
	var j int = 10 
	for ; j > 0; { //there are no while in Go, so u must user for, actualy no diff, but wierd 
		fmt.Println("in for used like while j =", j)
		j--
	}

	array := [...]string{"a", "b", "c", "d"}
	for num, val := range array { //there are foreach that can give you value num
		fmt.Printf("foreach array num, val: %v, %#v\n", num, val)
	}

	if 1 > 0 || 9 < 0 { //if works same
		fmt.Println("if 1 > 0 || 9 < 0 worked")
	}

	if 1 < 0 {
	} else {
		fmt.Println("if 1 < 0 else come in else")
	}

	a := 1

	switch (a) {
		case 0: fmt.Println("bruh")
		case 1:
			a++
			fmt.Println("in switch i increased value so now it:", a)
	}

	{ //you can actualy use blocks to enclose variables within
		b := "biba boba"
		fmt.Println("a in block:", b)
	}
	b := 10 // so i can change type
	fmt.Println("a out of block", b)
}