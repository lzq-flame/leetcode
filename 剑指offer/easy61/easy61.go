/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 20:05
 */

package main

func isStraight(nums []int) bool {
	hash := map[int]bool{}
	max, min := 0, 14
	for _, value := range nums {
		if value == 0 {
			continue
		}
		max = Max(max, value)
		min = Min(min, value)
		if hash[value] {
			return false
		}
		hash[value] = true
	}
	return max-min < 5
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{4, 7, 5, 9, 2}
	isStraight(nums)
}
