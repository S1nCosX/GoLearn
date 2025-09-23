package file

import (
	"fmt"
	"io"
	"os"
)

func File() {
	// file work is not this hard as it firstly looks like
	// lets create file
	fileName := "ThisFileCreatedByCode.txt"
	file, err := os.Create(fileName) // yep that simple.
	// important note: this is relative path, relative to directory where was program executed
	if err != nil {
		println(err.Error())
		return
	}
	defer os.Remove(fileName) // lets add into stack delete func for our file to not forget this
	for i := 0; i < 10; i++ {
		fmt.Fprintln(file, i) // this is simpliest way to write values into file.
		// F is kinda prefix means File, there Print(ln|f) funcs
	}
	file.Close() // close file

	file, _ = os.Open(fileName) // like with create, but this func simply open file
	defer file.Close()
	for {
		var val int
		n, err := fmt.Fscanln(file, &val)
		if err == io.EOF {
			print(err.Error())
			break
		}
		println("readed ", n, "bytes from file and value", val)
	}
}
