/**
2 * @Author: lzq
3 * @Date: 2021/10/17 10:31
4 */

//最长公共子序列
//https://leetcode-cn.com/problems/longest-common-subsequence/
package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	length1, length2 := len(text1), len(text2)
	dp := make([][]int, length1+1)
	for i := 0; i < length1+1; i++ {
		dp[i] = make([]int, length2+1)
	}
	for i := 0; i < length1+1; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < length2+1; i++ {
		dp[0][i] = 0
	}
	for i := 1; i < length1+1; i++ {
		for j := 1; j < length2+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[length1][length2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	subsequence := longestCommonSubsequence("bsbininm", "jmjkbkjkv")
	fmt.Println(subsequence)

}
