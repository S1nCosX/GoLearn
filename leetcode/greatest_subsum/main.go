package main

import "fmt"

func maxSubArray(nums []int) []int {
    subsumArrays := make([]int, 0)
    for i := 0; i < len(nums);{
        subsum := 0
        sign := nums[i] >= 0
        for ;i < len(nums) && ((sign && nums[i] >= 0) || (!sign && nums[i] <0)); i++ {
            subsum += nums[i]
        }

		for ;sign && len(subsumArrays) >= 2; {
			prevPositive := subsumArrays[len(subsumArrays) - 2]
			prevNegative := subsumArrays[len(subsumArrays) - 1]
			fmt.Println(prevPositive, prevNegative)
			
			if (prevNegative + prevPositive > 0) {
				subsum += prevNegative + prevPositive
				subsumArrays = subsumArrays[:len(subsumArrays) - 2]
			} else {
				break
			}
		}
		subsumArrays = append(subsumArrays, subsum)
    }
	var maximum int = -1e9
	if (len(subsumArrays) == 1 && subsumArrays[0] <= 0) {
		for _, val := range(nums) {
			if maximum < v {
				maximum = v
			}
		}
		return maximum
	}
	else {
		for _, val := range(subsumArrays) {
			if maximum < v {
				maximum = v
			}
		}
		return maximum
	}
	return subsumArrays
}

func main () {
	var array = [...]int{1, 2, 3, -4, -5, 6, 7 ,8, -9, 10}
	
	slices := maxSubArray(array[:])
	fmt.Println(slices)
}