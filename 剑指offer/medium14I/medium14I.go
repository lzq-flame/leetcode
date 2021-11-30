/**
 * @Author: liuzhiqi
 * @Data: 2021/11/26 14:24
 */

package main

import "fmt"

func cuttingRope(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		currMax := 0
		for j := 1; j <= i/2; j++ {
			currMax = max(currMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = currMax
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(cuttingRope(10))
}
