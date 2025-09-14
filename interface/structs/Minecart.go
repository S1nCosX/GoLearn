package structs

import "fmt"

type Minecart struct {
	Speed int
	Items []string
}

func (this Minecart) Draw() {
	fmt.Println("|     |\n\\_____/\n[]  []")
}

func (this Minecart) Move(time int) int {
	return this.Speed * time * 3
}
