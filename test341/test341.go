package main

import (
	"fmt"
	"reflect"
)

type NestedInteger struct {
	value interface{}
	next  *NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	if _, ok := this.value.(int); ok {
		return true
	} else {
		return false
	}
}

func (this NestedInteger) GetInteger() int {

	return int(reflect.ValueOf(this.value).Int())
}

func (n *NestedInteger) SetInteger(value int) {
	if n.IsInteger() == true {
		reflect.ValueOf(n.value).Elem().SetInt(int64(value))
	}
}

func (this *NestedInteger) Add(elem NestedInteger) {
	if this == nil {
		this = &elem
	}
	for this.next != nil {
		this = this.next
	}
	this.next = &elem
}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	val []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var vals []int
	vals = dfs(nestedList, vals)
	return &NestedIterator{vals}

}
func dfs(nestedList []*NestedInteger, val []int) []int {
	for _, ni := range nestedList {
		if ni.IsInteger() {
			val = append(val, ni.GetInteger())

		} else {
			val = dfs(ni.GetList(), val)
		}
	}
	return val
}

func (this *NestedIterator) Next() int {
	val := this.val[0]
	this.val = this.val[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	if len(this.val) > 0 {
		return true
	} else {
		return false
	}
}

func main() {
	// 定义切片。
	slice := []int{0, 1, 2, 3, 4}
	var test1 func()
	test1 = func() {
		fmt.Printf("%p \n", slice) // 0xc000072030 与实参的地址相同
		slice[0] = 111             // 修改形参会影响到实参

		// 但是 append()添加数据时并不会影响实参(的长度)。 (可以通过返回值覆盖(修改)实参的值(长度)；或者通过切片的指针修改实参)
		slice = append(slice, 66, 77, 88)
		fmt.Printf("%p \n", slice)
	}
	test1()
	fmt.Printf("%p \n", slice) // 0xc000072030

	fmt.Println(slice) // [111 1 2 3 4]  (修改形参会影响到实参)
}

// 切片作为函数的参数（地址传递） (切片变量名本身就是一个地址)
func test(s []int) {
	fmt.Printf("%p \n", s) // 0xc000072030 与实参的地址相同
	s[0] = 111             // 修改形参会影响到实参

	// 但是 append()添加数据时并不会影响实参(的长度)。 (可以通过返回值覆盖(修改)实参的值(长度)；或者通过切片的指针修改实参)
	s = append(s, 66, 77, 88)
	fmt.Printf("%p \n", s)
}
