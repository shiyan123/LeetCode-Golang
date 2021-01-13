package main

import (
	"fmt"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
123
321
*/

func main() {
	a := 1534236469
	fmt.Println(reverse(a))
}

//func reverse(x int) int {
//	INT_MAX := 2147483647
//	INT_MIN := -2147483648
//
//	if x == 0 {
//		return 0
//	}
//	var result int
//	str := strconv.Itoa(x)
//	temp := make([]string, len(str))
//	for k, v := range str {
//		temp[len(str)-1-k] += string(v)
//	}
//	str = strings.Join(temp, "")
//	if str[len(str)-1:] == "-" {
//		result, _ = strconv.Atoi(str[:len(str)-1])
//		if result > INT_MAX || result < INT_MIN {
//			return 0
//		}
//		return 0 - result
//	} else {
//		result, _ = strconv.Atoi(str)
//		if result > INT_MAX || result < INT_MIN {
//			return 0
//		}
//		return result
//	}
//}

func reverse(x int) int {
	y := 0
	for x != 0 {
		if y > 214748364 || y < -214748364 {
			return 0
		}
		y = y*10 + x%10
		x = x / 10
	}
	return y
}
