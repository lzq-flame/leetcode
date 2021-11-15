/**
2 * @Author: lzq
3 * @Date: 2021/10/18 10:22
4 */

//缺失的第一个正数
//https://leetcode-cn.com/problems/first-missing-positive/
package main

func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}

func main() {

}
