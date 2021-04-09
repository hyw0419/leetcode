package main

import "fmt"

//做法不太好 还是快慢指针，遍历一次
func main() {
	removeDuplicates([]int{1, 1, 2, 2, 2, 3})
}
func removeDuplicates(nums []int) int {
	l := len(nums)
	i := 0
	res := l
	for i < l-1 {
		t := i + 1
		if nums[i] == nums[i+1] {
			j := i + 1
			for j < l-1 && nums[j] == nums[j+1] {
				j++
			}
			t = j
			copy(nums[i+1:], nums[t:])
		}
		res = res - (t - i - 1)
		l = res
		fmt.Println(res)
		i++
	}
	fmt.Print(nums)
	fmt.Println(res)
	return 0
}
