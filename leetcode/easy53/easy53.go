package main

import "fmt"

//最大连续子数组和
//https://leetcode-cn.com/problems/maximum-subarray/

func maxSubArrayT(nums []int) int {
	result, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return result
}

//动态规划
func maxSubArray(nums []int) int {
	n := len(nums)
	ans := nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(dp[i], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	array := maxSubArray(nums)
	fmt.Println(array)
}
