package main

import "fmt"

func containsDuplicate(nums []int) bool {
	duplicateMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if duplicateMap[num] {
			return duplicateMap[num]
		}
		duplicateMap[num] = true
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}
