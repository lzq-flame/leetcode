/**
2 * @Author: lzq
3 * @Date: 2021/10/24 15:27
4 */

//子集
//https://leetcode-cn.com/problems/subsets/
package main

func subsets(nums []int) [][]int {

	ans := make([][]int, 0)
	if nums == nil {
		return ans
	}
	set := []int{}
	var dfs func(curr int)
	dfs = func(curr int) {
		if curr == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[curr])
		dfs(curr + 1)
		set = set[:len(set)-1]
		dfs(curr + 1)
	}
	dfs(0)
	return ans
}
