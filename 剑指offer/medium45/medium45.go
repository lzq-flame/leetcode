/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 19:23
 */

package main

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y < sx*y+x
	})
	ans := []byte{}
	for _, value := range nums {
		ans = append(ans, strconv.Itoa(value)...)
	}
	return string(ans)
}
