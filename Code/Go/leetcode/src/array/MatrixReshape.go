package array


func matrixReshape(nums [][]int, r int, c int) [][]int {
	// nums为空 or 
	if len(nums)==0 || len(nums)*len(nums[0]) != r*c {
		return nums
	}
	var result [][]int
	var res []int
	var l = 0
	for i:=0; i<len(nums); i++ {
		for j:=0; j < len(nums[0]); j++ {
			res = append(res, nums[i][j])
			l++
			if l % c == 0 {
				result = append(result, res)
				res = []int{}
				l = 0
			}
		}
	}
	return result
}
