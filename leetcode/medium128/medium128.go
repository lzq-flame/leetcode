/**
 * @Author: liuzhiqi
 * @Data: 2021/11/10 09:07
 */

//最长连续序列
//https://leetcode-cn.com/problems/longest-consecutive-sequence/
package main

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentStreak := 1
			currentNum := num
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

func main() {

}
