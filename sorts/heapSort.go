package main

import "fmt"

func moveHeap(arr []int, start, end int) {
	dad := start
	// 找到第一个子节点
	son := start * 2 + 1
	for son < end && arr[son] <= arr[son+1] {
		son++
	}
	if arr[dad] < arr[son] {
		arr[dad], arr[son] = arr[son], arr[dad]
	}
	return
}
// 建立大顶堆，即升序排列，复杂度Onlogn，不稳定
func buildHeap(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}
	// 找到最后一个父节点
	for i:=length/2-1; i>=0; i-- {
		// 移动堆，使其变成大顶堆
		// 循环完成后，根节点将成为数列的最大值
		moveHeap(arr, i, length-1)
	}
	// 弹出最大值，将其移动到数列尾部
	arr[0], arr[length-1] = arr[length-1], arr[0]
	// 对未完成的n-1的数列递归执行
	// 每次递归都会将其时数组中的最大值移动到尾部
	buildHeap(arr[0:length-1])
	return arr
}

func main() {
	arr := []int{3, 1, 7, 9, 8, 6, 3}
	fmt.Println(buildHeap(arr))
}