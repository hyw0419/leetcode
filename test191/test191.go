package main

import "fmt"

func main() {
	fmt.Print(hammingWeight(4))
}

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res = res + int((num & 1))
		num = num >> 1
	}

	return res
}
