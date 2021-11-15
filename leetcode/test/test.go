package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Sort interface {
	Sort(nums []int, len int)
}

type QuickSort struct {
	nums []int
}

func partition(nums []int, left, right int) int {
	i, j, pivot := left, right, nums[left]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return j
}

func randomizedPartition(nums []int, left, right int) int {
	i := getRandomNumber(right-left+1) + 1
	nums[i], nums[right] = nums[right], nums[i]
	return partition(nums, left, right)
}

func (quickSort *QuickSort) randomizedQuickSort(nums []int, left, right int) {
	if left < right {
		pos := partition(nums, left, right)
		quickSort.randomizedQuickSort(nums, left, pos-1)
		quickSort.randomizedQuickSort(nums, pos+1, right)
	}
}

func getRandomNumber(size int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(size)
}

func main() {
	arr := make([]int, 0)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr = append(arr, 5)
	arr = append(arr, 6)
	arr1 := make([]int, 6)
	//arr1 = append(arr1, 1, 2, 3, 4, 5, 6)
	copy(arr[0:2], arr1[3:])
	for _, value := range arr {
		fmt.Println(value)
	}
	fmt.Println("=============")
	fmt.Println(len(arr))

}
