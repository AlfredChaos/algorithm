package main

import "fmt"

func binary(arr []int, num int) int {
	length := len(arr)
	start := 0
	end := length - 1
	for start <= end {
		middle := (start + end) / 2
		if arr[middle] == num {
			return middle
		} else if arr[middle] > num {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return 0
}

func main() {
	arr := []int{1, 3, 3, 6, 7, 8, 9}
	fmt.Println(binary(arr, 9))
}