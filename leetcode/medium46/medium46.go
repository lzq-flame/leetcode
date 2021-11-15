/**
2 * @Author: lzq
3 * @Date: 2021/9/27 11:03
4 */

//全排列
//https://leetcode-cn.com/problems/permutations/
package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	length := len(nums)
	ans := make([][]int, 0)
	if length == 0 || length == 1 {
		ans = append(ans, nums)
	}
	used := make([]bool, length)
	path := make([]int, 0)
	DFS(length, 0, used, nums, path, &ans)
	return ans
}

func DFS(length, depth int, used []bool, nums, path []int, ans *[][]int) {
	if depth == length {
		current := make([]int, len(path))
		copy(current, path)
		*ans = append(*ans, current)
		return
	}
	for i := 0; i < length; i++ {
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			DFS(length, depth+1, used, nums, path, ans)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}

func main() {
	nums := []int{5, 4, 6, 2}
	//i := permute(nums)
	//
	//fmt.Println(i)
	//fmt.Println(nums)
	//f(nums)
	//fmt.Println(nums)
	ans := make([][]int, 0)
	ans = append(ans, nums)
	fmt.Println(ans)
	nums[0] = 0
	fmt.Println(ans)

}
