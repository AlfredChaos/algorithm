package main

import "fmt"

// https://leetcode.cn/problems/happy-number/
// 解决这道题有两个思路
// 1、找出凑成快乐数的数学特征
// 2、找出无限循环的边界
// 通过举例，18、17、15，发现
// 最后总会遇到循环的n，所以只要找环即可
func isHappy(n int) bool {
	record := make(map[int]bool)
	record[n] = true
	for n != 1 {
		sum := 0
		for n != 0 {
			x := n % 10
			n = n / 10
			sum += x * x
		}
		if _, ok := record[sum]; ok {
			return false
		}
		record[sum] = true
		n = sum
	}
	return true
}
