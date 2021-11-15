/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 19:36
 */

package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrices := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		if max < prices[i]-minPrices {
			max = prices[i] - minPrices
		}
		if minPrices > prices[i] {
			minPrices = prices[i]
		}
	}
	return max
}
