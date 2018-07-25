package array

// TODO: 因为go的原因，没有处理溢出，程序是有问题的
// https://leetcode-cn.com/problems/reverse-integer
func reverse(x int) int {
	tx := x  // 剩下没有处理的部分
	rem := 0 //余数，也即最低位数字
	res := 0 //结果
	for tx != 0 {
		rem = tx % 10
		tx /= 10
		res = res*10 + rem
	}
	t := pow(2, 31)
	if res < (-t) || res > (t-1) {
		res = 0
	}
	return res
}

func pow(x int, n int) int {
	ans := 1

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}
