package main

import (
	"fmt"
)

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49

输入：height = [4,3,2,1,4]
输出：16

*/

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	if len(height) == 2 {
		if height[0] > height[1] {
			return 1 * height[1]
		}
		return 1 * height[0]
	}
	x, y, area := 0, 0, 0
	i, j := 0, len(height)-1 //双指针
	for j-i != 0 {
		y = j - i
		if height[i] > height[j] {
			x = height[j]
			j -= 1
		} else {
			x = height[i]
			i += 1
		}
		if x*y > area {
			area = x * y
		}
	}
	return area
}
