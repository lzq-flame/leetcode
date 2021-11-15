/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 21:53
 */

//把数字翻译成字符串
package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	s := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(s); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := s[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}

func main() {
	temp, err := strconv.Atoi("05")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)
}
