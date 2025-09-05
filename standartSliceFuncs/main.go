package main

import (
	"fmt"
	"sort"
	"reflect"
)

func SortLib() {
	// there are exclusively types sorts like Ints, Float64s, Strings, its simply calls sort.Sort(slice) as its writen here https://pkg.go.dev/sort#example-Float64s
	// also fuctrions that check is array sorted, and search functions
	floatSlice := []float64{4, 2, 3, 1, 5, 8}
	fmt.Println("slice:", floatSlice, "is sorted func result for this slice:", sort.Float64sAreSorted(floatSlice))
	sort.Float64s(floatSlice) //lets sort slice
	fmt.Println("slice:", floatSlice, "is sorted func result for this slice:", sort.Float64sAreSorted(floatSlice))
	fmt.Println("slice:", floatSlice, "search(3) result:", sort.SearchFloat64s(floatSlice, 3))

	// if you want to sort random slice, you can use Slice func (slice, less func(i, j int))
	sliceOfAnonymStruct := []struct {
		int
		string
	} {{1, "2"}, {3, "4"}, {5, "6"}, {5, "9"}}
	fmt.Println("slice of struct before sort", sliceOfAnonymStruct)
	sort.Slice(sliceOfAnonymStruct, func(i, j int) bool {
			if sliceOfAnonymStruct[i].int == sliceOfAnonymStruct[j].int {
				return sliceOfAnonymStruct[i].string > sliceOfAnonymStruct[j].string
			}
			return sliceOfAnonymStruct[i].int > sliceOfAnonymStruct[j].int
		})
	fmt.Println("slice of struct after sort", sliceOfAnonymStruct)

	// also there is Sort(array Interface), Interface means "array" must have Len() int, Swap(i, j int) and Less(i, j int) functions
	// i prepared struct and Interface interpretation for this structs slice
	sliceOfStruct := SomeContainerStruct{[]SomeStruct{{1, "2"}, {3, "4"}, {5, "9"}, {5, "6"}}}
	fmt.Println("slice of struct before sort", sliceOfStruct)
	sort.Sort(sliceOfStruct)
	fmt.Println("slice of struct after sort", sliceOfStruct)
	// also for idk reason here is Reverse func
	sort.Sort(sort.Reverse(sliceOfStruct))
	// after some attempts i understand that it's returns struct with less function like return !original.Less()
	fmt.Println("reversed sorted slice:", sliceOfStruct)
}

func main() {
	SortLib()
	// for equal used reflect.DeepEqual
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{1, 2, 3, 4}
	// fmt.Println("slice1 == slice2:", slice1 == slice2) WRONG... normaly slices can be compared only with nil
	fmt.Println("reflect.DeepEqual(slice1, slice2):", reflect.DeepEqual(slice1, slice2))
}