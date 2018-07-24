package array

// [存在重复元素](https://leetcode-cn.com/problems/contains-duplicate/description/)
func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for _, num := range nums {
		if _, ok := mp[num]; ok {
			return true
		}
		mp[num] = true
	}
	return false
}
