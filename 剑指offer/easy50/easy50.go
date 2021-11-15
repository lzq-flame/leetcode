/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 15:13
 */

//第一个只出现一次的字符
package main

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	hash := map[byte]int{}
	for _, value := range s {
		hash[byte(value)]++
	}
	for _, value := range s {
		if hash[byte(value)] == 1 {
			return byte(value)
		}
	}
	return ' '
}
