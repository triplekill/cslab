package array

// 杨辉三角
// https://leetcode-cn.com/problems/pascals-triangle/description/
func generate(numRows int) [][]int {
	// o line-1 为1，其他位置根据相应的值填充
	var result [][]int
	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j <= i; j++ {
			tmp := 1
			if j==0 || j == i {
				//
			} else {
				tmp = result[i-1][j-1]+result[i-1][j]
			}
			row = append(row, tmp)
		}
		result = append(result, row)
	}

	return result
}
