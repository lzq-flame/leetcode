/**
2 * @Author: lzq
3 * @Date: 2021/10/11 16:58
4 */

package medium31

//下一个排列
//https://leetcode-cn.com/problems/next-permutation/
func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := length - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}
