package main

import "fmt"

/*
输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.
*/

func main() {
	str := "III"
	fmt.Println(romanToInt(str))
}

func romanToInt(s string) int {
	result := 0
	tempIntMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	for i := 0; i < len(s); i++ {
		j := 0
		if i+1 < len(s) {
			j = i + 1
		}
		key := ""
		if j != 0 {
			key = string(s[i]) + string(s[j])
		}
		var flag bool
		if temp, ok := tempIntMap[key]; ok {
			flag = true
			result += temp
		} else {
			if temp, ok := tempIntMap[string(s[i])]; ok {
				result += temp
			}
		}
		if flag {
			i = i + 1
		}
	}
	return result
}
