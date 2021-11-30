/**
 * @Author: liuzhiqi
 * @Data: 2021/11/19 15:52
 */

package main

func verifyPostorder(postorder []int) bool {
	return check(postorder, 0, len(postorder)-1)
}

func check(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	mid := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && check(postorder, i, mid-1) && check(postorder, mid, j-1)
}
