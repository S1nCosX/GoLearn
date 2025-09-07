package main

import (
	"importedModule/somePackage" // for linking package we must use comman 'go mod edit -replace="importedModule/somePackage=[PATH]"' our situation is 'go mod edit -replace="importedModule/somePackage=../importModule"' and go mod tidy to clearify requirements
	"fmt"
)

func main() {
	fmt.Println(somePackage.Hehe())
}