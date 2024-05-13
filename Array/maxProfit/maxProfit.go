// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
package main

import "math"

func maxProfit(prices []int) int {
	max := 0
	minprice := math.MaxInt
	profit := 0
	for i := 0; i < len(prices); i++ {
		profit = prices[i] - minprice
		if profit > max {
			max = profit
		}
		if minprice > prices[i] {
			minprice = prices[i]
		}
	}
	return max
}

func main() {

}
