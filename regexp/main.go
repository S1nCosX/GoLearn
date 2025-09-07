package main

import (
	"regexp"
	"fmt"
)

func main() {
	regexp := regexp.MustCompile("Boba+")
	fmt.Println(string(regexp.Find([]byte("amogus big Boooobaaaa Bobaaa"))))
}