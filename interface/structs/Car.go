package structs

import "fmt"

type Car struct {
	Speed int
	Name  string
	Seats [4]string
}

func (this Car) Draw() {
	fmt.Println("------\n|    \\\n[]----[]")
}

func (this Car) Move(time int) int {
	return this.Speed * time
}
