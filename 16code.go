package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个包括n 个整数的数组nums和 一个目标值target。找出nums中的三个整数，使得它们的和与target最接近。返回这三个数的和。假定每组输入只存在唯一答案。

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

*/
func main() {
	nums := []int{-1, 0, 1, 1, 55}
	target := 3
	fmt.Println(threeSumClosest(nums, target))
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	if len(nums) <= 3 {
		return sum(nums)
	}
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return ans
			}
		}
	}
	return ans
}

func sum(nums []int) (total int) {
	for _, v := range nums {
		total += v
	}
	return total
}
