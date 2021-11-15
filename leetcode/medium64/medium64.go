/**
2 * @Author: lzq
3 * @Date: 2021/10/30 15:47
4 */

//最小路径和
//https://leetcode-cn.com/problems/minimum-path-sum/
package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

func minPathSum1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([]int, columns)

	dp[0] = grid[0][0]
	for j := 1; j < columns; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if j == 0 {
				dp[j] = dp[0] + grid[i][0]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[columns-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	sum := minPathSum1(nums)
	fmt.Println(sum)
}
