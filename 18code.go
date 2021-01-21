package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组
*/

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	sum := func(nums []int) (total int) {
		for _, v := range nums {
			total += v
		}
		return
	}
	if len(nums) < 4 {
		return [][]int{}
	}
	if len(nums) == 4 {
		if sum(nums) != target {
			return [][]int{}
		}
		return [][]int{nums}
	}
	sort.Ints(nums)
	fmt.Println(nums)
	result := make([][]int, 0)
	for i := 0; i <= len(nums)-4; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j <= len(nums)-3; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				s := sum([]int{nums[i], nums[j], nums[left], nums[right]})
				if s == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[right] == nums[right-1] { //右指针移相同值跳过
						right--
					}
					right--
					for left < right && nums[left] == nums[left+1] { //左指针移相同值跳过
						left++
					}
					left++
				} else if s > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}
