package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; (j-i <= k) && (j < length); j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
