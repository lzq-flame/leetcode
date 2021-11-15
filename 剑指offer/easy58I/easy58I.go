/**
 * @Author: liuzhiqi
 * @Data: 2021/11/15 13:55
 */

package main

import "strings"

func reverseWords(s string) string {
	s = process(s)
	temp := strings.Split(s, " ")
	for i, n := 0, len(temp)-1; i < n; i, n = i+1, n-1 {
		temp[i], temp[n] = temp[n], temp[i]
	}
	return strings.Join(temp, " ")
}
func process(s string) string {
	var res []byte
	l := len(s) - 1
	for l >= 0 && s[l] == ' ' {
		l--
	}
	flag := 1
	for i := 0; i <= l; i++ {
		if s[i] != ' ' {
			res = append(res, s[i])
			flag = 0
		}
		if flag == 0 && s[i] == ' ' {
			res = append(res, s[i])
			flag = 1
		}
	}
	return string(res)
}
