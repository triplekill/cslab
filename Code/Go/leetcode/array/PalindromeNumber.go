package array

// 回文数 https://leetcode-cn.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	remain := x
	rev := 0
	for remain != 0 {
		rev = rev*10 + remain%10
		remain /= 10
	}
	return rev == x
}

func isPalindrome2(abc int) bool {
	if abc >= 0 && abc <= 9 {
		return true
	}
	if abc < 0 || abc%10 == 0 {
		return false
	}
	x := abc
	r := 0
	for ; x >= 10; x = x / 10 {
		r += x % 10
		r *= 10
	}
	return abc == x+r
}
