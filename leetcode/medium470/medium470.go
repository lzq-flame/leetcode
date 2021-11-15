/**
2 * @Author: lzq
3 * @Date: 2021/10/18 17:04
4 */

//https://leetcode-cn.com/problems/implement-rand10-using-rand7/
//Rand7() 实现 Rand10()
package main

func Rand10() int {
	for {
		num := (Rand7()-1)*7 + Rand7()
		if num <= 40 {
			return num%10 + 1
		}
		num = (num-40-1)*7 + Rand7()
		if num <= 60 {
			return num%10 + 1
		}
		num = (num-60-1)*7 + Rand7()
		if num <= 20 {
			return num%10 + 1
		}
	}
}

func Rand7() int {
	return 0
}
