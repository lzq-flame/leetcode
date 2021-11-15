/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 10:48
 */

//替换空格
package main

func replaceSpace(s string) string {
	ans := ""
	for _, value := range s {
		if value == ' ' {
			ans += "%20"
		} else {
			ans += string(value)
		}
	}
	return ans
}
