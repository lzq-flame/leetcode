/**
2 * @Author: lzq
3 * @Date: 2021/10/27 16:36
4 */

package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ans := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ans[i+j+1] += x * y
		}
	}
	for i := m + n - 1; i > 0; i-- {
		ans[i-1] += ans[i] / 10
		ans[i] = ans[i] % 10
	}
	k := 1
	s := ""
	if ans[0] != 0 {
		k = 0
	}
	for ; k < m+n; k++ {
		s += strconv.Itoa(ans[k])
	}
	return s
}

func main() {
	i := multiply("123", "456")
	fmt.Println(i)
	fmt.Println()
}
