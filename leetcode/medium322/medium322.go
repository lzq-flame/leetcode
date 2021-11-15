/**
2 * @Author: lzq
3 * @Date: 2021/10/25 11:06
4 */

package main

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = max
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
