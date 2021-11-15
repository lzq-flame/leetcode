/**
2 * @Author: lzq
3 * @Date: 2021/9/8 10:38
4 */

//https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

//第k个最大的数

package main

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := heapSize - 1; i >= heapSize-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize>>1 - 1; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	left, right, largest := i>>1+1, i>>1+2, i
	for left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	for right < heapSize && nums[right] > nums[largest] {
		largest = right
	}
	if i != largest {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, heapSize)
	}
}

func main() {

}
