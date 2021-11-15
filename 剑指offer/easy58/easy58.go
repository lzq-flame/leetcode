/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 11:00
 */

//左旋转字符串
package main

func reverseLeftWords(s string, n int) string {
	if len(s) == 0 {
		return ""
	}
	return s[n:] + s[:n]
}
