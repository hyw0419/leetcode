package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1}
	fmt.Print(trap(nums))
}
func trap(height []int) int {
	i := 0
	j := len(height) - 1
	sum := 0
	for i < j {
		//左边找到第一个不是0
		for i < len(height) && height[i] == 0 {
			i++
		}
		//右边找到第一个不是0的数
		for j >= 0 && height[j] == 0 {
			j--
		}
		for k := i; k <= j; k++ {
			if height[k] == 0 {
				sum++
			} else {
				height[k]--
			}
		}
	}
	return sum
}
