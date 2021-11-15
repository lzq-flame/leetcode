/**
2 * @Author: lzq
3 * @Date: 2021/9/13 10:16
4 */

package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/two-sum/
//两数之和

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		if key, ok := hashTable[a]; ok {
			return []int{i, key}
		}
		hashTable[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	sum := twoSum(nums, 9)
	fmt.Println(sum)
}
