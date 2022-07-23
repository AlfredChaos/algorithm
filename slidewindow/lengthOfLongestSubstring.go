package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	duplicate := make(map[byte]bool)
	left := 0
	right := 0
	max := 0
	for right < len(s) {
		if _, ok := duplicate[s[right]]; ok {
			length := right - left
			if length > max {
				max = length
			}
			delete(duplicate, s[left])
			left++
			continue
		}
		duplicate[s[right]] = true
		right++
	}
	if max < len(duplicate) {
		max = len(duplicate)
	}
	return max
}

func main() {
	s := "au"
	fmt.Println(lengthOfLongestSubstring(s))
}

