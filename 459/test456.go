package main

import "fmt"

func main() {

	var nums = []int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print(find132pattern(nums))
}
func find132pattern(nums []int) bool {
	for i := 1; i < len(nums)-1; i++ {

		j := i + 1
		max := nums[j]
		if nums[i] > nums[j] {
			for ; j < len(nums); j++ {
				if nums[i] > nums[j] && nums[j] > max {
					max = nums[j]
				}
			}
			for k := 0; k < i; k++ {
				if nums[k] < max && nums[i] > max {
					return true
				}
			}
		}

	}
	return false
}
