/**
2 * @Author: lzq
3 * @Date: 2021/10/24 14:00
4 */

//最小栈
//https://leetcode-cn.com/problems/min-stack/
package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{minStack: []int{math.MaxInt64}, stack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(val, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
