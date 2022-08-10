package main

// https://leetcode.cn/problems/array-partition/
// 贪心算法，只要每次找局部最优解即可，但局部最优解不一定就是全局最优解
// 贪心失效，则需要动态规划
func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    result := 0
    for i:=0; i<len(nums); i+=2 {
        result += nums[i]
    }
    return result
}