package main

import "fmt"

func bubble(arr []int) []int {
	// 每一个数跟其它所有数比较，冒泡
	length := len(arr)
	// 两层循环，On^2，稳定
	// 小的数值从后面冒上来
	for i:=0; i<length-1; i++ {
		for j:=length-1; j>i; j-- {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 3, 7, 9, 8, 6, 3}
	fmt.Println(bubble(arr))
}