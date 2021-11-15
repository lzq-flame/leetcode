/**
2 * @Author: lzq
3 * @Date: 2021/11/2 21:58
4 */

//组合总数
//https://leetcode-cn.com/problems/combination-sum/
package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(nums []int, Target, sum, startIndex int)

	dfs = func(nums []int, Target, sum, startIndex int) {
		if sum > Target {
			return
		}
		if sum == Target {
			B := make([]int, len(path))
			copy(B, path)
			ans = append(ans, B)
			return
		}
		for i := startIndex; i < len(nums); i++ {
			sum += nums[i]
			path = append(path, nums[i])
			dfs(nums, Target, sum, i)
			sum -= nums[i]
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, target, 0, 0)
	return ans
}

func main() {
	nums := []int{2, 3, 7, 6}
	sum := combinationSum(nums, 7)
	fmt.Println(sum)
}
