package advances

import (
	"fmt"
	"interface/interfaces"
)

func PrintString(inter interface{}) {
	fmt.Println(inter)
}

func AdvancesDemo() {
	// also we can use empty interface to make func input abstract
	PrintString(123)
	PrintString("asdsad")
	// also we can check is struct realisation contains interface
	var someStruct interface{} = struct {
		int
		string
	}{1, "biba"}
	_, ok := someStruct.(interfaces.Movable)
	fmt.Println(ok)
	// also there are construction named Type switch:
	switch val := someStruct.(type) {
	case float32:
		fmt.Println("val is float", val)
		break
	case bool:
		fmt.Println("val is a bool", val)
		break
	default:
		fmt.Println("omg this is something else")
	}
}
