package main

type SomeContainerStruct struct {
	container []SomeStruct
}

func (this SomeContainerStruct) Len() int{
	return len(this.container)
}

func (this SomeContainerStruct) Cap() int{
	return cap(this.container)
}

func (this SomeContainerStruct) Less(i, j int) bool{
	return this.container[i].Less(this.container[j])
}

func (this SomeContainerStruct) Swap(i, j int) {
	buffer := this.container[i]
	this.container[i] = this.container[j]
	this.container[j] = buffer
}