package array

// 成最多的水 https://leetcode-cn.com/problems/container-with-most-water/description/
func maxArea(height []int) int {
	maxar, l, r := 0, 0, len(height)-1
	for l < r {
		maxar = max(maxar, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxar
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
