package main

import "fmt"

func insertion(arr []int) []int {
	length := len(arr)
	var prev, current int
	// On^2, 稳定, 插入排序
	// 假定第一个元素已排序，从后往前比较
	// 前者需要往后移位，腾出空间插入
	for i:=1; i<length; i++ {
		prev = i-1
		current = arr[i]
		for prev >= 0 && arr[prev] > current {
			arr[prev+1] = arr[prev]
			prev--
		}
		arr[prev+1] = current
	}
	return arr
}

func main() {
	arr := []int{3, 1, 7, 9, 8, 6, 3}
	fmt.Println(insertion(arr))
}