package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

//快速排序
func quickSort(nums []int, left, right int) {
	if left < right {
		mid := partition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func partition(nums []int, left, right int) int {
	i, j, pivot := left, right, nums[left]
	for i != j {
		for i < j && nums[j] > pivot {
			j--
		}
		if i < j {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		for i < j && nums[i] < pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return j
}

//堆排序
func heapSort(nums []int) {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := heapSize - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
}

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

//归并排序
func mergeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right) >> 1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	i, j, k := left, mid+1, 0
	B := make([]int, right-left+1)
	for i <= mid && j <= right {
		if nums[i] > nums[j] {
			B[k] = nums[j]
			j++
		} else {
			B[k] = nums[i]
			i++
		}
		k++
	}
	for i <= mid {
		B[k] = nums[i]
		k++
		i++
	}
	for j <= right {
		B[k] = nums[j]
		k++
		j++
	}
	copy(nums[left:right+1], B)
}

//冒泡排序
func bubbleSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

//插入排序
func selectionSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	length := len(nums)
	var minIndex int
	for i := 0; i < length-1; i++ {
		minIndex = i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

//希尔排序
func shellSort(nums []int) {
	length := len(nums)
	//外层步长控制
	for step := length / 2; step >= 1; step /= 2 {
		//开始插入排序
		for i := step; i < length; i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}

		}
	}
}

//计数排序
func countingSort(nums []int, maxValue int) {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	sortedIndex := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		bucket[nums[i]] += 1
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			nums[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
}

//桶排序
func bucketSort(nums []int) {
	min, max := math.MaxInt32, math.MinInt32

	//计算最大值与最小值
	for i := 0; i < len(nums); i++ {
		max = Max(max, nums[i])
		min = Min(min, nums[i])
	}

	//计算桶的数量
	bucketNum := (max-min)/len(nums) + 1
	bucketArr := make([][]int, 0)
	for i := 0; i < bucketNum; i++ {
		bucketArr = append(bucketArr, make([]int, 0))
	}

	//将每个元素放入桶
	for i := 0; i < len(nums); i++ {
		num := (nums[i] - min) / len(nums)
		bucketArr[num] = append(bucketArr[num], nums[i])
	}

	//对每个桶进行排序
	for i := 0; i < len(bucketArr); i++ {
		sort.Ints(bucketArr[i])
	}
	//将桶中的元素赋值到原序列
	index := 0
	for i := 0; i < len(bucketArr); i++ {
		for j := 0; j < len(bucketArr[i]); j++ {
			nums[index] = bucketArr[i][j]
			index++
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func radixSort(nums []int) {
	//存数组中最大的数字，为了知道循环几次
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		max = Max(max, nums[i])
	}

	//计算最大数是几位数
	max1 := strconv.Itoa(max)
	maxLength := len(max1)
	fmt.Println(maxLength)

	//用于临时存储数据的数组
	var temp [10][]int

	for i, _ := range temp {
		temp[i] = make([]int, len(nums))
	}

	//用于存储桶内的元素位置
	counts := make([]int, 10)

	//循环的次数
	n := 1
	for i := 0; i < maxLength; i++ {
		for j := 0; j < len(nums); j++ {
			//计算余数
			remainder := (nums[j] / n) % 10
			// 把便利店数据放在指定数组中，有两个信息，放在第几个桶+数据应该放在第几位
			temp[remainder][counts[remainder]] = nums[j]
			//记录数量
			counts[remainder]++
		}

		//记录取的数字应该放到的位置
		index := 0
		//每一轮循环之后把数字取出来
		for k := 0; k < len(counts); k++ {
			//记录数量的数组中当前余数记录不为零
			if counts[k] != 0 {
				for l := 0; l < counts[k]; l++ {
					//取出元素
					nums[index] = temp[k][l]
					index++
				}
				//取出后把数量置为零
				counts[k] = 0
			}
		}

		n *= 10
	}

}
func main() {
	nums := []int{5, 2, 3, 1}
	selectionSort(nums)
	fmt.Println(nums)
}
