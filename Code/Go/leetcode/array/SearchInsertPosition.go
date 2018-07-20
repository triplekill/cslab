package array

// [搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/description/)
func searchInsert(nums []int, target int) int {
	// 二分查找
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			return mid
		}
	}
	return low
}
