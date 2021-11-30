/**
2 * @Author: lzq
3 * @Date: 2021/10/25 16:52
4 */

package main

//旋转图像
//https://leetcode-cn.com/problems/rotate-image/
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
