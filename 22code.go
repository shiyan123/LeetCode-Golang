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
	n := 2
	fmt.Println(generateParenthesis(n))
}

var (
	List = []string{}
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	dfs(n, "")
	return List
}

func dfs(n int, str string) {
	if len(str) == 2*n {
		if check(str) {
			List = append(List, str)
			return
		}
		return
	}
	dfs(n, str+"(")
	dfs(n, str+")")
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
