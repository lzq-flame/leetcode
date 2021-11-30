/**
 * @Author: liuzhiqi
 * @Data: 2021/11/17 18:50
 */

package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" || len(digits) == 0 {
		return []string{}
	}
	hash := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ans := make([]string, 0)
	path := make([]byte, 0)
	var backTrace func(s string, index1 int)
	backTrace = func(s string, index1 int) {
		if index1 >= len(s) {
			ans = append(ans, string(append([]byte(nil), path...)))
			return
		}
		temp := hash[s[index1]]
		for j := 0; j < len(temp); j++ {
			path = append(path, temp[j])
			backTrace(s, index1+1)
			path = path[:len(path)-1]
		}

	}
	backTrace(digits, 0)
	return ans
}

func main() {
	s := "23"
	combinations := letterCombinations(s)
	fmt.Println(combinations)
}
