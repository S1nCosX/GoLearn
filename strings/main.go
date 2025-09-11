package main

import (
	"fmt"
	"strs/stringLib"
)

func main() {
	// lets talk about strings
	// strings strings strings
	// strings in Go are constant. end.
	str := "Some text"
	// str[0] = "a"
	// that line will trigger trouble: 'cannot assign to str[0] (neither addressable nor a map index expression)'
	// so this is read-only type
	fmt.Println(str)
	// stings are read only for 1 reason, i think. Optimization. Strings are using only bytes that are have meaning
	// For example: english a = [97], len("a") = 1, ф = [209 132], len("ф") = 2, byte can hold value up to 255
	fmt.Printf("dec a=%v, len(a)=%d\n", []byte("a"), len("a"))
	fmt.Printf("dec ф=%v, len(ф)=%d\n", []byte("ф"), len("ф"))
	// so we want to iterate on string with foreach
	helloWorld := "Привет, мир"
	fmt.Printf("lets wrinte %%v, in golang it'll write go type\n")
	fmt.Printf("right way to iterate give us this: ")
	for _, val := range helloWorld {
		fmt.Printf("%c ", val)
	}
	// cause if you want to iterate on string with item number. you'll iterate on bytes and get bullshit
	fmt.Printf("\nin wrong way: ")
	for i := 0; i < len(helloWorld); i++ {
		fmt.Printf("%c ", helloWorld[i])
	}
	fmt.Println()

	// So this is moment where we discovers runes. It's Golangs unicode interpritation, there all symbols have equal size, that give us oportunity to change it. For example:
	someStr := []rune("some string")
	fmt.Println("runes converted to str", string(someStr))
	someStr[4] = '♂'
	fmt.Println("runes converted to str changed symbol", string(someStr))
	// also this is kinda method to change strings:
	str = "BOSS OF THE GYM"
	fmt.Println("str before editing:", str)
	strToRunes := []rune(str)
	strToRunes = strToRunes[:4]
	str = string(strToRunes)
	fmt.Println("str after editing:", str)
	fmt.Println("also, lets talk about strings lib\n----------------------\n")
	stringLib.LetsTalkBoutStringLibs()
}