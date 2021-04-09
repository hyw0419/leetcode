package main

import "fmt"

func main() {
	merge([]int{0, 2, 3, 2}, 4, []int{2, 2, 5}, 3)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	res := make([]int, 0, n+m)

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
		//fmt.Print(res)
	}
	if i >= m {
		res = append(res, nums2[j:]...)
	} else {
		res = append(res, nums1[i:]...)
	}
	fmt.Print(nums1)
	fmt.Print(res)
	copy(nums1[2:], res)

	fmt.Print(nums1)
}
