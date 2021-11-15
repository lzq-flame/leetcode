/**
2 * @Author: lzq
3 * @Date: 2021/10/3 15:32
4 */
//最长递增子序列
//https://leetcode-cn.com/problems/longest-increasing-subsequence/
package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxans := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxans = max(maxans, dp[i])
	}
	return maxans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums)+1)
	length := 1
	d[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[length] {
			length++
			d[length] = nums[i]
		} else {
			left, right, pos := 1, length, 0
			for left <= right {
				mid := (left + right) >> 1
				if d[mid] < nums[i] {
					pos = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return length
}

func main() {
	//nums := []int{7,7,7,7,7,7,7}
	//lis := lengthOfLIS(nums)
	//fmt.Println(lis)
}
