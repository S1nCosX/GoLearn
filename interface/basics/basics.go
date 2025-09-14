package basics

import (
	"fmt"
	"interface/interfaces"
	"interface/structs"
)

func TransportInfo(transport interfaces.Transport) {
	fmt.Println(transport.Move(1))
	transport.Draw()
}

func BasicsDemo() {
	// interfaces in go is thing is an requirement for class realisation
	// but instead of direct requirement in class this is requirement in func
	// also interfaces can be complex, for example in my code Transport is
	// combination of Drawable and Movable interfaces
	car := structs.Car{Speed: 123, Name: "Toyota", Seats: [4]string{"Gary", "Vova", "Anton", "Mr.Worldwide"}}
	minecart := structs.Minecart{Speed: 123, Items: []string{"ingot", "picaxe"}}
	TransportInfo(car)
	TransportInfo(minecart)
}
