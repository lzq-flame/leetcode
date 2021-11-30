/**
 * @Author: liuzhiqi
 * @Data: 2021/11/17 16:46
 */

package medium11

func maxArea(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		if height[left] < height[right] {
			ans = max(ans, (right-left)*height[left])
			left++
		} else {
			ans = max(ans, (right-left)*height[right])
			right--
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
