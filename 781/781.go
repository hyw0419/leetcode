package main

import "fmt"

func main() {
	fmt.Print(numRabbits([]int{0, 1, 1, 1, 0}))
}

func numRabbits(answers []int) int {
	check := make(map[int]int)
	extra := 0
	for _, v := range answers {
		//存在值跳过
		_, ok := check[v]
		if ok && check[v] > 0 {
			check[v] = v - 1
		} else {
			check[v] = v
			//fmt.Println(v)
			extra = extra + v + 1
		}
	}
	//fmt.Println(extra)
	return extra

}
