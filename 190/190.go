package main

import (
	"fmt"
)

func main() {
	fmt.Print(reverseBits(uint32(7)))
}

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res = res << 1
		res = res + uint32(num&1)
		num = num >> 1
	}

	return res
}
