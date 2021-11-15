/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 22:45
 */

package main

import "fmt"

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = max(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	value := maxValue(nums)
	fmt.Println(value)
}
