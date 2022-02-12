package main

import "fmt"

func quick(arr []int) {
	// 快排，Onlogn, 最差On^2，不稳定
	length := len(arr)
	if length <= 1 {
		return
	}
	left := 0
	right := length - 1
	// 以第一个元素为基准值
	key := arr[left]

	for left != right {
		// 移动指针
		for left < right && arr[right] >= key {
			right--
		}
		// 比基准值key小的数值，转到key的左边
		arr[left], arr[right] = arr[right], arr[left]
		// 移动指针
		for left < right && arr[left] <= key {
			left++
		}
		// 比基准值key大的数值，转到key的右边
		arr[left], arr[right] = arr[right], arr[left]
	}
	// 最终left的位置最终会落在基准值上
	// 因为left最终会落在第一个大于等于基准值的数值上
	// 而基准值左边的数都会小于基准值，所以left最终停在基准值本身
	// 如果存在等于基准值的数，它会排列在基准值左右，对结果来说并无影响

	quick(arr[0:left])
	// 注意基准值右边的数列，不包括基准值
	quick(arr[left+1:])
}

func main() {
	arr := []int{3, 1, 3, 9, 8, 6, 7}
	quick(arr)
	fmt.Println(arr)
}