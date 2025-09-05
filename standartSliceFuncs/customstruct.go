package main

type SomeStruct struct {
	fieldA int
	fieldB string
}

func (this SomeStruct) Less(other SomeStruct) bool {
	if this.fieldA == other.fieldA {
		return this.fieldB > other.fieldB
	}
	return this.fieldA > other.fieldA
}
