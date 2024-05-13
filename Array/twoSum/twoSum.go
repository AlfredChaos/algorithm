package main

import "fmt"

// https://leetcode.cn/problems/two-sum/
// 暴力枚举
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumHash(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for index, value := range nums {
		if i, ok := hashTable[target-value]; ok {
			return []int{i, index}
		}
		hashTable[value] = index
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
