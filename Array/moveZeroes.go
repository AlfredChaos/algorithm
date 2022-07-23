package main

import "fmt"

// https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	x := 0
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	for y := 0; y < len(nums); y++ {
		if nums[y] != 0 {
			if nums[x] == 0 {
				nums[x] = nums[y]
				nums[y] = 0
			}
			x++
		}
	}
	return
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
