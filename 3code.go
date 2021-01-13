package main

import "fmt"

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//示例 4:
//
//输入: s = ""
//输出: 0

func main() {
	str := "abcaecbb"
	result := lengthOfLongestSubstring(str)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]uint8
	left, right := 0, 0
	result := 0
	//左侧 开始查询
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for left < len(s) {
		//查到结尾
		if right < len(s) && bitSet[s[right]] == 0 {
			//不重复
			bitSet[s[right]] = 1
			right++
		} else {
			bitSet[s[left]] = 0
			left++
		}
		result = max(result, right-left)
	}
	return result
}
