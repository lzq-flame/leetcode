/**
 * @Author: liuzhiqi
 * @Data: 2021/11/29 15:03
 */

package main

import (
	"math"
)

func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	length := len(str)
	i := 0
	flag := true
	ans := 0
	for i < length && str[i] == ' ' {
		i++
	}
	if i < length && str[i] == '-' {
		flag = false
	}
	if i < length && (str[i] == '-' || str[i] == '+') {
		i++
	}
	for i < length && int(str[i]-'0') >= 0 && int(str[i]-'0') <= 9 {
		r := int(str[i] - '0')
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && r > 7) {
			if flag {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		ans = ans*10 + r
		i++
	}
	if flag {
		return ans
	} else {
		return -1 * ans
	}
}
