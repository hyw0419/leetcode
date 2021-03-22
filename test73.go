package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 0}, {1, 11, 1}, {111, 11, 11}}
	setZeroes1(matrix)
}
func setZeroes1(matrix [][]int) {
	x := len(matrix[0])
	y := len(matrix)
	i := 0
	j := 0
	resulx := make([]int, x)
	resulty := make([]int, y)
	for i < x {
		for j < y {
			if matrix[i][j] == 0 {
				resulty[j] = 1
				resulx[i] = 1
			}
			j++
		}
		j = 0
		i++
	}
	i = 0
	j = 0
	for i < x {
		for j < y {
			if resulx[i] == 1 || resulty[j] == 1 {
				matrix[i][j] = 0
			}
			j++
		}
		j = 0
		i++
	}
	fmt.Print(resulty)
	fmt.Print(resulx)
	fmt.Print(matrix)
}
