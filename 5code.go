package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

*/
func main() {
	s := "bbabc"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if s == "" || len(s) < 2 {
		return s
	}
	length, start, end, maxLen := len(s), 0, 0, 1
	dp := [1000][1000]bool{}
	for r := 1; r < length; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > maxLen {
					maxLen = r - l + 1
					start = l
					end = r
				}
			}
		}
	}
	return s[start : end+1]
}

//func longestPalindrome(s string) string {
//	var index = 0
//	var max = 0
//	var size = len(s)
//	var res string
//	for index < size {
//		left := index
//		right := index
//		for left >= 0 && s[left] == s[index] { // 如果中心点左边和中心点相同，则向左边扩散检测
//			left--
//		}
//		for right < size && s[right] == s[index] { // 如果中心点右边和中心点相同，则向右边扩散
//			right++
//		}
//		next := right                                          // 下一个点向右移动
//		for left >= 0 && right < size && s[left] == s[right] { // 如果左右相等，则向两边扩散
//			left--
//			right++
//		}
//		index = next
//		midMaxLen := right - left + 1
//		if midMaxLen > max {
//			res = s[left+1 : right]
//			max = midMaxLen
//		}
//	}
//	return res
//}
