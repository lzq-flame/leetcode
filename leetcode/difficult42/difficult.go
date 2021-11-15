/**
2 * @Author: lzq
3 * @Date: 2021/10/4 9:02
4 */

//接雨水
//https://leetcode-cn.com/problems/trapping-rain-water/
package main

import "fmt"

type stack struct {
	stack []int
	size  int
}

func (s *stack) pop() int {
	a := s.stack[s.size-1]
	s.stack = s.stack[:s.size-1]
	s.size--
	return a
}

func (s *stack) push(a int) {
	s.stack = append(s.stack, a)
	s.size++
}

func (s *stack) Size() int {
	return s.size
}

func Stack() *stack {
	return &stack{stack: make([]int, 0), size: 0}
}

func (s *stack) top() int {
	if !s.empty() {
		return s.stack[s.size-1]
	}
	return 0
}

func (s *stack) empty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	s := Stack()
	ans, current := 0, 0
	for current < len(height) {
		for s.empty() && height[current] > height[s.top()] {
			top := s.pop()
			if s.empty() {
				break
			}
			distance := current - s.top() - 1
			bondHeight := min(height[current], height[s.top()]) - height[top]
			ans += distance * bondHeight
		}
		s.push(current)
		current++
	}
	return ans

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 4, 2, 1, 2, 1}
	i := trap(height)
	fmt.Println(i)
}
