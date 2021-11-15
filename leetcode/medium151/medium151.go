/**
2 * @Author: lzq
3 * @Date: 2021/10/7 16:43
4 */

//翻转字符串里的单词
//https://leetcode-cn.com/problems/reverse-words-in-a-string/
package main

import "fmt"

type Stack struct {
	stack []interface{}
	size  int
}

func newStack() *Stack {
	return &Stack{stack: make([]interface{}, 0), size: 0}
}

func (s *Stack) push(a interface{}) {
	s.stack = append(s.stack, a)
	s.size++
}
func (s *Stack) pop() interface{} {
	a := s.stack[s.size-1]
	s.stack = s.stack[:s.size-1]
	s.size--
	return a
}
func (s *Stack) empty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func reverseWords(s string) string {
	stack := newStack()
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else {
			j := i + 1
			for j < len(s) && s[j] != ' ' {
				j++
			}
			stack.push(s[i:j])
			i = j
		}
	}
	for !stack.empty() {
		pop := stack.pop()
		ans += pop.(string)
		ans += " "
	}
	ans = ans[:len(ans)-1]
	return ans
}

func main() {
	a := "the sky is blue"
	words := reverseWords(a)
	fmt.Println(words)
}
