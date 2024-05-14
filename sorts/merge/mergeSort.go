package main

import "fmt"

func mergeSort(arr []int) []int {
	// 归并排序， Onlogn, 稳定
	// mergeSort是递归切割每段
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	// 对于单独一段，进行举例证明，总结规律
	result := make([]int, 0)
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

func main() {
	arr := []int{3, 1, 7, 9, 8, 6, 3}
	fmt.Println(mergeSort(arr))
}