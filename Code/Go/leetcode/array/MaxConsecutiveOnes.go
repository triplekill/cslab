package array

// https://leetcode-cn.com/problems/max-consecutive-ones/description/
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	sum := 0
	for _, num := range nums {
		if num == 1 {
			sum++
			if sum > max {
				max = sum
			}
		} else {
			sum = 0
		}
	}
	return max
}
