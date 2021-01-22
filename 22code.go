package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

输入：n = 1
输出：["()"]

*/

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}

	var dfs func(n int, str string)
	var list []string

	dfs = func(n int, str string) {
		if len(str) == 2*n {
			if check(str) {
				list = append(list, str)
				return
			}
			return
		}
		dfs(n, str+"(")
		dfs(n, str+")")
	}
	dfs(n, "")
	return list
}

func check(str string) bool {
	x := 0
	for k, v := range str {
		if string(v) != "(" && k == 0 {
			return false
		}
		if string(v) == "(" {
			x++
		} else {
			x--
		}
		if x < 0 {
			return false
		}
	}
	return x == 0
}

/*
func generateParenthesis(n int) []string {
	return generateHelper(n, 0, "", nil)
}

func generateHelper(nLeft, nRight int, str string, result []string) []string {
	if nLeft<=0 && nRight <=0 {
		result = append(result, str)
		return result
	}

	if nLeft > 0 {
		result = generateHelper(nLeft-1, nRight+1, str+"(", result)
	}
	if nRight > 0 {
		result = generateHelper(nLeft, nRight-1, str+")", result)
	}

	return result
}
*/
