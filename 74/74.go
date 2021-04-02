package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {

	j := len(matrix) - 1
	i := 0
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		} else {
			return true
		}
	}
	return false
}
