package array

func majorityElement(nums []int) int {
	mp := make(map[int]int)
	l := len(nums)
	var result int
	mid := l / 2
	for _, i := range nums {
		mp[i]++
		if mp[i] > mid {
			result = i
			return result
		}
	}
	return result
}
