/**
2 * @Author: lzq
3 * @Date: 2021/10/6 21:43
4 */

//字符串转换整数atoi
//https://leetcode-cn.com/problems/string-to-integer-atoi/
package main

import (
	"fmt"
	"math"
)

//0-9的ascil码分别是48-57，"-"是45，"+"是43
func myAtoi(s string) int {
	length := len(s)
	sum := 0
	flag := true
	i := 0
	for i < length && s[i] == ' ' {
		i++
	}
	if i < length && s[i] == '-' {
		flag = false
	}
	if i < length && (s[i] == '-' || s[i] == '+') {
		i++
	}
	for i < length && int(s[i]-'0') >= 0 && int(s[i]-'0') <= 9 {
		r := int(s[i] - '0')
		if sum > math.MaxInt32/10 || (sum == math.MaxInt32/10 && r > 7) {
			if flag {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		sum = sum*10 + r
		i++
	}
	if flag {
		return sum
	} else {
		return -1 * sum
	}
}

func main() {
	atoi := myAtoi("123")
	fmt.Println(atoi)
}
