/**
 * @Author: liuzhiqi
 * @Data: 2021/11/9 17:42
 */

package main

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {

}
