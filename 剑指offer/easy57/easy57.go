/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 11:21
 */

package main

func twoSum(nums []int, target int) []int {
	hash := map[int]bool{}
	for _, value := range nums {
		if hash[target-value] {
			return []int{value, target - value}
		}
		hash[value] = true
	}
	return []int{}
}
