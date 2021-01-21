package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串

class Solution:
    def isValid(self, s):
        while '{}' in s or '()' in s or '[]' in s:
            s = s.replace('{}', '')
            s = s.replace('[]', '')
            s = s.replace('()', '')
        return s == ''

*/

func main() {
	str := "([])[]{}"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	var stack []string
	for i := 0; i < n; i++ {
		fmt.Println(stack)
		if temp, ok := pairs[string(s[i])]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != temp {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, string(s[i]))
		}
	}
	return len(stack) == 0
}
