package main

import "fmt"

func selection(arr []int) []int {
	var index int
	length := len(arr)
	// On^2，稳定, 选择排序（我的实现是稳定的）
	// 跟冒泡排序的不同是，每次需要从后面的数列中寻找最小值
	for i:=0; i<=length-1; i++ {
		index = i
		for j:=i+1; j<=length-1; j++ {
			if arr[j] < arr[index] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}

func main() {
	arr := []int{3, 1, 7, 9, 8, 6, 3}
	fmt.Println(selection(arr))
}