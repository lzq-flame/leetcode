/**
2 * @Author: lzq
3 * @Date: 2021/10/31 14:58
4 */

package main

import (
	"fmt"
	"math"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[left]
	}
	for right-left > 1 {
		mid := (left + right) >> 1
		if nums[mid] > nums[left] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[right]
}

func average(nums []float64) float64 {
	var x float64
	x = 0
	for _, value := range nums {
		x += value
	}
	return x / float64(len(nums))
}

func fc(nums []float64) float64 {
	var x float64
	x = 0
	average := average(nums)

	for _, value := range nums {
		x += math.Pow(value-average, 2)
	}
	return math.Sqrt(x / float64(len(nums)))
}

func xs(nums []float64) float64 {
	var x float64
	x = 0
	average := average(nums)

	for _, value := range nums {
		x += abs(value - average)
	}
	return x / float64(len(nums))
}

func abs(a float64) float64 {
	if a < 0 {
		return -1 * a
	}
	return a
}

func zscore(nums []float64) {
	average := average(nums)
	fc := fc(nums)
	for _, value := range nums {
		fmt.Println(value, ":", (value-average)/fc)
	}
}

func main() {
	//nums := []float64{23, 23, 27, 27, 39, 41, 47, 49, 50, 52, 54, 54, 56, 57, 58, 58, 60, 61}
	fat := []float64{9.5, 26.5, 7.8, 17.8, 31.4, 25.9, 27.4, 27.2, 31.2, 34.6, 42.5, 28.8, 33.4, 30.2, 34.1, 32.9, 41.2, 35.7}
	zscore(fat)
	fmt.Println(average(fat))
	fmt.Println(fc(fat))
}
