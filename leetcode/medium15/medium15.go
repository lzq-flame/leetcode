/**
2 * @Author: lzq
3 * @Date: 2021/9/13 16:24
4 */
//三数之和
//https://leetcode-cn.com/problems/3sum/
package main

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		target := -1 * nums[i]
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}

		}
	}
	return ans
}

func main() {

}
