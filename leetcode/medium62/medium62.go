/**
2 * @Author: lzq
3 * @Date: 2021/10/31 14:23
4 */

package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//只用一个一位数组保存dp迭代结果
func uniquePaths1(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				continue
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[n-1]
}
func main() {
	i := uniquePaths1(3, 7)
	fmt.Println(i)
}
