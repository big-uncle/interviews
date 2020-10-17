package main

import "fmt"

// 一个数组成V字排序：开始递减，然后递增，然后输出最小值的索引
//利用二分法  O(n)  时间复杂度

func main() {
	a := []int{7, 6, 5, 2}
	fmt.Println(Findmix(a))
}

func Findmix(nums []int) int {
	x, y := 0, len(nums)-1
	res := 0
	for x <= y {
		if nums[x] > nums[x+1] && nums[y] > nums[y-1] {
			x++
			y--
		} else {
			if nums[x] < nums[y] {
				res = x
			} else {
				res = y
			}
			break
		}

	}
	return res
}
