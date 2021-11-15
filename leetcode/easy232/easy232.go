/**
2 * @Author: lzq
3 * @Date: 2021/10/4 10:19
4 */
//用栈实现队列
//https://leetcode-cn.com/problems/implement-queue-using-stacks/
package main

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

type MyQueue struct {
	stack1, stack2 *stack
}

func Constructor() MyQueue {
	return MyQueue{stack1: Stack(), stack2: Stack()}
}

func (this *MyQueue) Push(x int) {
	this.stack1.push(x)
}

func (this *MyQueue) Pop() int {
	if this.stack2.empty() {
		for !this.stack1.empty() {
			this.stack2.push(this.stack1.pop())
		}
	}
	return this.stack2.pop()

}

func (this *MyQueue) Peek() int {
	if this.stack2.empty() {
		for !this.stack1.empty() {
			this.stack2.push(this.stack1.pop())
		}
	}
	return this.stack2.top()
}

func (this *MyQueue) Empty() bool {
	if this.stack1.empty() && this.stack2.empty() {
		return true
	}
	return false
}
