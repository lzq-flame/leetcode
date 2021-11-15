/**
2 * @Author: lzq
3 * @Date: 2021/10/15 10:36
4 */

//编辑距离
//https://leetcode-cn.com/problems/edit-distance/
package main

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}
	dp := make([][]int, n+1)
	for key, _ := range dp {
		dp[key] = make([]int, m+1)
	}

	//边界状态初始化
	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			left_down := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				left_down++
			}
			dp[i][j] = min(left, min(down, left_down))
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
