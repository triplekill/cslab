package array

// 需要考虑 [3,3] 6的情况
// https://leetcode-cn.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	res := []int{}
	for index, num := range nums {
		mp[num] = index
	}
	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		if _, ok := mp[comp]; ok && mp[comp] != i {
			res = []int{i, mp[comp]}
			break
		}
	}
	return res
}
