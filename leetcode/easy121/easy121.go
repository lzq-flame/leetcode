/**
2 * @Author: lzq
3 * @Date: 2021/9/17 16:32
4 */

package main

import (
	"fmt"
)

//买股票的最佳时机
//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	max := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if max < prices[i]-minPrice {
			max = prices[i] - minPrice
		}
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}
	return max
}

func main() {
	p := []int{7, 1, 5, 3, 6}
	fmt.Println(maxProfit(p))
}
