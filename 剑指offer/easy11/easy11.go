/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 14:51
 */

//旋转数组中的最小数字
package main

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) >> 1
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right = right - 1
		}
	}
	return numbers[left]
}
