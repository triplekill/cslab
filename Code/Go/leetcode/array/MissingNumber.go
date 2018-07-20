package array

func missingNumber(nums []int) int {
	mp := make(map[int]bool)
	max := 0
	for _, num := range nums {
		mp[num] = true
		if num > max {
			max = num
		}
	}
	for i := 0; i <= max+1; i++ {
		if !mp[i] {
			return i
		}
	}
	return 0
}
