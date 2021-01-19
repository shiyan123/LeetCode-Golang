package main

import "fmt"

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

func main() {
	digits := "2"
	fmt.Println(letterCombinations(digits))
}

var (
	NumMap = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	List = []string{}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	recursion(digits, 0, "")
	return List
}

func recursion(digits string, index int, str string) {
	if len(digits) == index {
		List = append(List, str)
	} else {
		temp := NumMap[string(digits[index])]
		for i := 0; i < len(temp); i++ {
			recursion(digits, index+1, str+temp[i])
		}
	}
}
