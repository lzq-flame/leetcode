/**
2 * @Author: lzq
3 * @Date: 2021/9/23 18:37
4 */

package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	add := 0
	result := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result = x + y + add
		add = result / 10
		ans = strconv.Itoa(result%10) + ans
	}

	return ans

}
