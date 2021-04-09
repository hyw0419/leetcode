package main

import "fmt"

func main() {
	fmt.Print(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

func search(nums []int, target int) bool {
	length := len(nums)
	left := length - 1
	right := 0
	for right <= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[right] && nums[mid] == nums[left] {
			right++
			left--
			continue
		}
		if nums[mid] < nums[left] {
			if nums[mid] < target && nums[left] > target {
				right = mid + 1
			} else {
				left = mid - 1
			}
		} else {
			if nums[mid] < target && nums[right] < target {
				right = mid + 1
			} else {
				left = mid - 1
			}
		}
	}
	return false
}
