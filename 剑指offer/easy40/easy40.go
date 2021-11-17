/**
 * @Author: liuzhiqi
 * @Data: 2021/11/16 10:49
 */

package main

import "fmt"

//构建大顶堆
func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize>>1 - 1; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	left, right, largest := i<<1+1, i<<1+2, i
	if left < heapSize && nums[largest] < nums[left] {
		largest = left
	}
	if right < heapSize && nums[largest] < nums[right] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		maxHeapify(nums, largest, heapSize)
	}
}

func heapSort(nums []int) {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := heapSize - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	length := len(arr)
	m := length - k
	buildMaxHeap(arr, length)
	for i := len(arr) - 1; i >= 0 && m >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		length--
		m--
		maxHeapify(arr, 0, length)
	}
	return arr[:length+1]
}

func main() {
	nums := []int{5, 6, 3, 1, 8, 10}
	numbers := getLeastNumbers(nums, 3)
	fmt.Println(numbers)
}
