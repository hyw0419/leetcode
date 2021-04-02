package main

import "fmt"

func main() {
	fmt.Print(clumsy(10))
}

func clumsy(N int) int {
	if N == 1 || N == 2 || N == 0 {
		return N
	}
	if N == 3 {
		return 6
	}
	n := N
	sum := n*(n-1)/(n-2) + n - 3
	n = n - 4
	fmt.Println(sum)
	for n >= 4 {
		sum = sum - n*(n-1)/(n-2) + n - 3
		n = n - 4
	}
	fmt.Println(sum)
	switch n {
	case 0:
		return sum
	case 1:
		return sum - 1
	case 2:
		return sum - 2
	case 3:
		return sum - 6
	}
	return sum
}
