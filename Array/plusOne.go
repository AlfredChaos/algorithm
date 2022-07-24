package main

import "fmt"

// https://leetcode.cn/problems/plus-one/
func plusOne(digits []int) []int {
	length := len(digits)
	digits[length-1] = digits[length-1] + 1
	if digits[length-1] >= 10 {
		digits[length-1] = 0
		if length-1 == 0 {
			return capArray(digits)
		}
		for i := length - 2; i >= 0; i-- {
			digits[i] = digits[i] + 1
			if digits[i] >= 10 {
				digits[i] = 0
				if i == 0 {
					return capArray(digits)
				}
				continue
			} else {
				break
			}
		}
	}
	return digits
}

func capArray(digits []int) []int {
	result := make([]int, 1)
	result[0] = 1
	result = append(result, digits...)
	return result
}

func main() {
	digits := []int{4, 3, 2, 9}
	fmt.Println(plusOne(digits))
}
