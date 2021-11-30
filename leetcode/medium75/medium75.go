/**
 * @Author: liuzhiqi
 * @Data: 2021/11/18 11:10
 */

package main

import "fmt"

func sortColors(nums []int) {
	if len(nums) == 1 {
		return
	}
	i := 0
	for n := 0; n < len(nums); n++ {
		if nums[n] == 0 {
			nums[i], nums[n] = nums[n], nums[i]
			i++
		}
	}
	for n := i; n < len(nums); n++ {
		if nums[n] == 1 {
			nums[n], nums[i] = nums[i], nums[n]
			i++
		}
	}
}

func sortColors1(nums []int) {
	zero := 0
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors1(nums)
	fmt.Println(nums)
}
