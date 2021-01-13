package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/

func main() {
	str := []string{"baab", "bacb", "b", "cbc"}
	fmt.Println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	for _, v := range strs {
		if len(v) == 0 {
			return ""
		}
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	str := ""
	for i := 0; i < minLen; i++ {
		flag := 0
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] == strs[j+1][i] {
				flag += 1
			} else {
				flag -= 1
			}
		}
		if flag == len(strs)-1 {
			str += string(strs[0][i])
		} else {
			break
		}
	}
	return str
}
