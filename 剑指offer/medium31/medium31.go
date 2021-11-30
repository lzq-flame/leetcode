/**
 * @Author: liuzhiqi
 * @Data: 2021/11/29 10:47
 */

package main

type Stack struct {
	stack []int
	size  int
}

func NewStack() *Stack {
	return &Stack{stack: make([]int, 0), size: 0}
}

func (s *Stack) push(i int) {
	s.stack = append(s.stack, i)
	s.size++
}
func (s *Stack) pop() int {
	a := s.stack[s.size-1]
	s.stack = s.stack[:s.size-1]
	s.size--
	return a
}

func (s *Stack) top() int {
	a := s.stack[s.size-1]
	return a
}

func (s *Stack) empty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack, j, N := []int{}, 0, len(pushed)
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) != 0 && j < N && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
	}
	return j == N
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	poped := []int{4, 3, 5, 1, 2}
	println(validateStackSequences(pushed, poped))
}
