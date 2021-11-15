/**
2 * @Author: lzq
3 * @Date: 2021/10/17 14:40
4 */

//复原IP地址
//https://leetcode-cn.com/problems/restore-ip-addresses/
package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	var dfs func(subRes []string, start int)
	dfs = func(subRes []string, start int) {
		if len(subRes) == 4 && start == len(s) {
			ans := strings.Join(res, ".")
			res = append(res, ans)
			return
		}
		if len(subRes) == 4 && start > len(s) {
			return
		}
		for length := 1; length <= 3; length++ {
			if start+length-1 >= len(s) {
				return
			}
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			if num, _ := strconv.Atoi(str); num > 255 {
				return
			}
			subRes = append(subRes, str)
			dfs(subRes, start)
			subRes = subRes[:len(subRes)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}
