package main

//import "fmt" // you can import packages like this
import (
	//	"fmt" // also like this if there more than 1 pakage
	. "fmt" // if you use dot before package name,you can use it without package prefix, like using namespace in C++
	_ "math" // if you want to use some lib in future but now you dont use it, you can add _ and compilier will not argue
	"testComplexProgram/absolutlyRandomPackage" // if you want to divide project into different packages, you can do it inside of module and call it like "module_name/package_path"
	// IMPORTANT NOTE: this is way inside of Module, not like package name inside of module
	// to get outsider package we use "go get [package_name]" it'll automaticaly added into module
	"rsc.io/quote" // import outsider spackage
)

func main () {
	Println("used fmt lib func Println without prefix")
	absolutlyRandomPackage.Aboba("ALABUGA POLITECH")

	Println(quote.Hello())
}