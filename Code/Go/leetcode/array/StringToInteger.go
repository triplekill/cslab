package array

import "math"

// 字符串转数字 https://leetcode-cn.com/problems/string-to-integer-atoi/description/
func myAtoi(str string) int {
	//init
	strLen, sign, rs, idx := len(str), 1, 0, 0
	//skip whitespaces
	for idx < strLen && str[idx] == ' ' {
		idx++
	}
	//+-
	if idx < strLen {
		if str[idx] == '-' {
			sign = -1
			idx++
		} else {
			if str[idx] == '+' {
				idx++
			}
		}
	}
	//digits
	for ; idx < strLen; idx++ {
		if str[idx] >= '0' && str[idx] <= '9' {
			rs = rs*10 + int(str[idx]-'0')
			if rs*sign > math.MaxInt32 {
				return math.MaxInt32
			}
			if rs*sign < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}

	return rs * sign
}

/*
func myAtoi(str string) int {
	res := 0
	start := 0
	neg := false
	overflow := false

	for start < len(str) && str[start] == ' ' {
		start++
	}
	if start < len(str) && (str[start] == '+' || str[start] == '-') {
		if str[start] == '-' {
			neg = true
		}
		start++
	}

	for ; start < len(str); start++ {
		if str[start] >= '0' && str[start] <= '9' {
			res = res*10 + int((str[start] - '0'))
			if res > math.MaxInt32 || res < math.MinInt32 {
				overflow = true
				break
			}
		} else {
			// 是字符串，则跳出循环
			break
		}
	}
	if neg {
		res *= -1
	}
	maxi := pow(2, 31)
	// > < 两个有点瞎搞
	if overflow && res > 0 {
		res = maxi - 1
	}
	if overflow && res < 0 {
		res = maxi * (-1)
	}
	return res
}
*/
// func pow(x int, n int) int {
// 	ans := 1

// 	for n != 0 {
// 		ans *= x
// 		n--
// 	}
// 	return ans
// }
