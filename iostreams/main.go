package main

import (
	file "iostreams/file"
	reader "iostreams/reader"
)

func main() {
	println("----------------------------\nReader")
	reader.Reader()
	println("----------------------------\nFile")
	file.File()
}
