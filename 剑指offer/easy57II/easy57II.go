/**
 * @Author: liuzhiqi
 * @Data: 2021/11/26 14:39
 */

package easy57II

func findContinuousSequence(target int) [][]int {
	left, right := 1, 2
	ans := make([][]int, 0)
	for left < right {
		sum := (left + right) * (right - left + 1) / 2
		if sum == target {
			res := make([]int, 0)
			for i := left; i <= right; i++ {
				res = append(res, i)
			}
			ans = append(ans, res)
			left++
		} else if target > sum {
			right++
		} else {
			left++
		}
	}
	return ans
}
